<template>
  <div class="md-layout">
    <div class="md-layout-item">
      <md-card>
        <md-card-header class="md-card-header-icon md-card-header-green">
          <md-card-header-text>
            <div class="card-icon" style="margin-left: -15px">
              <md-icon>devices</md-icon>
            </div>
            <h4 class="title">设备列表</h4>
          </md-card-header-text>
          <div class="button-container" style="margin-top: 10px">
            <md-button class="md-success md-sm" @click="toDeviceAdd">
              <md-icon class="md-size-1x" style="height: 0; width: 0">add</md-icon>
              添加设备
            </md-button>
          </div>
        </md-card-header>
        <md-card-content>
          <md-table
            :value="tableData"
            class="paginated-table table-striped table-hover"
          >
            <md-table-toolbar>
              <md-field>
                <label for="pages">每页:</label>
                <md-select v-model="pagination.perPage" name="pages">
                  <md-option
                    v-for="item in pagination.perPageOptions"
                    :key="item"
                    :label="item"
                    :value="item"
                  >
                    {{ item }}
                  </md-option>
                </md-select>
              </md-field>

              <md-field>
                <md-input
                  type="search"
                  class="mb-3"
                  clearable
                  style="width: 200px"
                  placeholder="搜索"
                  v-model="searchQuery"
                >
                </md-input>
              </md-field>
            </md-table-toolbar>

            <md-table-row slot="md-table-row" slot-scope="{ item }">
              <!-- <md-table-cell md-label="ID">{{ item.id }}</md-table-cell> -->
              <md-table-cell md-label="设备ID">
                {{item.device_id}}
                <span v-show="item.super" class="badge badge-info">管理员</span>
              </md-table-cell>
              <md-table-cell md-label="备注">
                {{substring(item.desc)}}
                <md-tooltip md-direction="bottom">{{item.desc}}</md-tooltip>
              </md-table-cell>
              <md-table-cell md-label="状态">{{ item.status }}</md-table-cell>
              <md-table-cell md-label="添加时间">{{ item.time }}</md-table-cell>
              <md-table-cell md-label="操作">
                <!-- <md-button
                  class="md-just-icon md-info md-simple"
                  @click.native="handleLike(item)"
                >
                  <md-icon>open_in_new</md-icon>
                  <md-tooltip md-direction="top">查看</md-tooltip>
                </md-button> -->
                <md-button
                  class="md-just-icon md-info md-simple"
                  @click.native="handleEdit(item)"
                >
                  <md-icon>dvr</md-icon>
                </md-button>
                &nbsp;&nbsp;&nbsp;&nbsp;
                <md-button
                  class="md-just-icon md-danger md-simple"
                  @click.native="handleDelete(item)"
                >
                  <md-icon>close</md-icon>
                </md-button>
              </md-table-cell>
            </md-table-row>
          </md-table>
          <div class="footer-table md-table">
            <table>
              <tfoot>
                <tr>
                  <th
                    v-for="item in footerTable"
                    :key="item.name"
                    class="md-table-head"
                  >
                    <div class="md-table-head-container md-ripple md-disabled">
                      <div class="md-table-head-label">
                        {{ item }}
                      </div>
                    </div>
                  </th>
                </tr>
              </tfoot>
            </table>
          </div>
        </md-card-content>
        <md-card-actions md-alignment="space-between">
          <div class="">
            <p class="card-category">
              展示 {{ from + 1 }} 到 {{ to }}, 共 {{ pagination.total }} 个
            </p>
          </div>
          <pagination
            class="pagination-no-border pagination-success"
            v-model="pagination.currentPage"
            :per-page="pagination.perPage"
            :total="pagination.total"
          >
          </pagination>
        </md-card-actions>
      </md-card>
      <modal v-if="deleteId" @close="closeDeleteModal">
        <template slot="header">
          <md-button class="md-simple md-just-icon md-round modal-default-button" @click="closeDeleteModal">
            <md-icon>clear</md-icon>
          </md-button>
        </template>

        <template slot="body">
          <p>确定要删除?</p>
        </template>

        <template slot="footer">
          <md-button class="md-simple" @click="closeDeleteModal">取消</md-button>
          <md-button class="md-success md-simple" @click="deleteUser">是的</md-button>
        </template>
      </modal>
    </div>
  </div>
</template>

<script>
import { Pagination, Modal } from "@/components";
import { deviceList, userDeviceList, deviceDelete } from "@/api/device";

export default {
  components: {
    Pagination,
    Modal
  },
  props: {
    userId: {
      type: String,
      default: ""
    },
    inUserDetail: {
      type: Boolean,
      default: false
    }
  },
  computed: {
    /** *
     * Returns a page from the searched data or the whole data. Search is performed in the watch section below
     */
    // queriedData() {
    //   let result = this.tableData;
    //   if (this.searchedData.length > 0) {
    //     result = this.searchedData;
    //   }
    //   return result.slice(this.from, this.to);
    // },
    to() {
      let highBound = this.from + this.pagination.perPage;
      if (this.pagination.total < highBound) {
        highBound = this.pagination.total;
      }
      return highBound;
    },
    from() {
      return this.pagination.perPage * (this.pagination.currentPage - 1);
    }
  },
  data() {
    return {
      pagination: {
        perPage: 10,
        currentPage: 1,
        perPageOptions: [5, 10, 25, 50],
        total: 0
      },
      footerTable: ["设备ID", "备注", "状态", "添加时间", "操作"],
      searchQuery: "",
      tableData: [],
      deleteId: ""
    };
  },
  created() {
    this.fecthDate();
  },
  methods: {
    fecthDate() {
      if (this.userId !== "") {
        userDeviceList(this.userId, { page: this.pagination.currentPage, limit: this.pagination.perPage }).then(res => {
          this.tableData = res.data.items;
          this.pagination.total = res.data.count;
        }).catch(err => {
          if (err.code === 404) {
            this.tableData = [];
          }
        });
      } else {
        if (!this.inUserDetail) {
          deviceList({ page: this.pagination.currentPage, limit: this.pagination.perPage }).then(res => {
            this.tableData = res.data.items;
            this.pagination.total = res.data.count;
          });
        }
      }
    },
    toDeviceAdd() {
      if (this.userId !== "") {
        this.$router.push({ name: "添加设备(给用户)", params: { id: this.userId }});
      } else {
        this.$router.push({ name: "添加设备" });
      }
    },
    handleEdit(item) {
      this.$router.push({ name: "设备详情", params: { id: item.id }});
    },
    handleDelete(item) {
      this.deleteId = item.id;
    },
    closeDeleteModal() {
      this.deleteId = "";
    },
    deleteUser() {
      // this.closeDeleteModal();
      deviceDelete(this.deleteId).then(res => {
        // console.log(res);
      });
      this.closeDeleteModal();
      this.fecthDate();
    },
    substring(str) {
      if (str.length <= 10) {
        return str;
      } else {
        return str.substring(0, 10) + "...";
      }
    }
  },
  watch: {
    /**
     * Searches through the table data by a given query.
     * NOTE: If you have a lot of data, it's recommended to do the search on the Server Side and only display the results here.
     * @param value of the query
     */
    searchQuery(value) {
      // let result = this.tableData;
      // if (value !== "") {
      //   result = this.fuseSearch.search(this.searchQuery);
      // }
      // this.searchedData = result;
    },
    "pagination.perPage"() {
      this.fecthDate();
    },
    "pagination.currentPage"() {
      this.fecthDate();
    },
    userId() {
      this.fecthDate();
    }
  }
};
</script>

<style lang="css" scoped>
.md-card .md-card-actions{
  border: 0;
  margin-left: 20px;
  margin-right: 20px;
}
</style>
