const getters = {
  isLoggedIn: state => !!state.user.token
};

export default getters;
