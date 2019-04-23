package emqx

import (
	"errors"
	"fmt"
)

type Banned struct {
	Who    string `json:"who"`
	As     string `json:"as"`
	Reason string `json:"reason"`
	Desc   string `json:"desc"`
	Until  int    `json:"until"`
}

func Banneds(page, limit int) ([]Banned, Meta, error) {
	var result struct {
		Code    int      `json:"code"`
		Message string   `json:"message"`
		Data    []Banned `json:"data"`
		Meta    Meta     `json:"meta"`
	}

	res, err := Rest.Get("/banned/").
		Query("_page", fmt.Sprintf("%v", page)).
		Query("_limit", fmt.Sprintf("%v", limit)).
		Send()
	if err != nil {
		return nil, Meta{}, err
	}

	err = res.Json(&result)
	if err != nil {
		return nil, Meta{}, err
	}

	if res.StatusCode() != 200 {
		return nil, Meta{}, fmt.Errorf("%s", res.StatusCode())
	}

	if result.Code > 0 {
		return nil, Meta{}, errors.New(result.Message)
	}

	return result.Data, result.Meta, nil
}

func CreateBanneds(banned *Banned) (Banned, error) {
	var result struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    Banned `json:"data"`
	}

	res, err := Rest.Post("/banned/").
		Json(banned).
		Send()
	if err != nil {
		return Banned{}, err
	}

	err = res.Json(&result)
	if err != nil {
		return Banned{}, err
	}

	if res.StatusCode() != 200 {
		return Banned{}, fmt.Errorf("%s", res.StatusCode())
	}

	if result.Code > 0 {
		return Banned{}, errors.New(result.Message)
	}

	return result.Data, nil
}

func DeleteBanneds(who, as string) error {
	var result struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	res, err := Rest.Delete(fmt.Sprintf("/banned/%v", who)).
		Query("as", as).
		Send()
	if err != nil {
		return err
	}

	err = res.Json(&result)
	if err != nil {
		return err
	}

	if res.StatusCode() != 200 {
		return fmt.Errorf("%s", res.StatusCode())
	}

	if result.Code > 0 {
		return errors.New(result.Message)
	}

	return nil
}
