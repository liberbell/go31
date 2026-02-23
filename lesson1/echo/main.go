package main

import (
	"net/http"
	"restapi/cache"
	"restapi/user"
	"strings"

	"github.com/asdine/storm"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func usersOptions(c echo.Context) error {
	methods := []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}
	c.Response().Header().Set("Allow", strings.Join(methods, ","))
	c.NoContent(http.StatusOK)
}

func userOptions(c echo.Context) error {
	methods := []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions}
	c.Response().Header().Set("Allow", strings.Join(methods, ","))
	c.NoContent(http.StatusOK)
}

func usersGetAll(c echo.Context) error {
	if cache.Serve(w, r) {
		return
	}

	users, err := user.All()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
		return
	}

	if c.Request().Method == http.MethodHead {
		return c.NoContent(http.StatusOK)
	}
	cw := cache.NewWriter(w, r)
	postBodyResponse(cw, http.StatusOK, jsonResponse{"users": users})
}

func usersPostOne(c echo.Context) error {
	u := new(user.User)
	err := bodyToUser(r, u)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
		return
	}
	u.ID = bson.NewObjectId()
	err = u.Save()
	if err != nil {
		if err == user.ErrRecordInvalid {
			return echo.NewHTTPError(http.StatusBadRequest)
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return
	}
	cache.Drop("/users")
	w.Header().Set("Location", "/users/"+u.ID.Hex())
	return c.NoContent(http.StatusCreated)
}

func usersGetOne(c echo.Context) error {
	if cache.Serve(w, r) {
		return
	}

	u, err := user.One(id)
	if err != nil {
		if err == storm.ErrNotFound {
			return echo.NewHTTPError(http.StatusNotFound)
			return
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
		return
	}
	if c.Request().Method == http.MethodHead {
		return c.NoContent(http.StatusOK)
	}
	cw := cache.NewWriter(w, r)
	postBodyResponse(cw, http.StatusOK, jsonResponse{"user": u})
}

func usersPutOne(c echo.Context) error {
	u := new(user.User)
	err := bodyToUser(r, u)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	u.ID = id
	err = u.Save()
	if err != nil {
		if err == user.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	cache.Drop("/users")
	cw := cache.NewWriter(w, r)
	postBodyResponse(cw, http.StatusOK, jsonResponse{"user": u})
}

func usersPatchOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	u, err := user.One(id)
	if err != nil {
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	err = bodyToUser(r, u)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}

	u.ID = id
	err = u.Save()
	if err != nil {
		if err == user.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	cache.Drop("/users")
	cw := cache.NewWriter(w, r)
	postBodyResponse(cw, http.StatusOK, jsonResponse{"user": u})
}

func usersDeleteOne(c echo.Context) error {
	err := user.Delete(id)
	if err != nil {
		if err == storm.ErrNotFound {
			return echo.NewHTTPError(http.StatusNotFound)
			return
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
		return
	}
	cache.Drop("/users")
	cache.Drop(cache.MakeResource(r))
	w.WriteHeader(http.StatusOK)
}

func root(c echo.Context) error {
	return c.String(http.StatusOK, "Running API v1")
}

func main() {
	e := echo.New()

	e.GET("/", root)
	u := e.Group("/users")
	u.OPTIONS("", usersOptions)
	u.HEAD("", usersGetAll)
	u.GET("", usersGetAll)
	u.POST("", usersPostOne)

	uid := u.Group("/id")
	uid.OPTIONS("", usersOptions)
	uid.HEAD("", usersGetOne)
	uid.GET("", usersGetOne)
	uid.PUT("", usersPutOne)
	uid.PATCH("", usersPatchOne)
	uid.DELETE("", usersDeleteOne)

	e.Start(":12345")
}
