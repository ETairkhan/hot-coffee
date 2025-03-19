package utils

import "log/slog"

func ReqGroup() slog.Attr {
	reqGroup := slog.Group(
		"request",
		"method", "GET",
	)
	return reqGroup
}

func PostGroup() slog.Attr {
	PostGroup := slog.Group(
		"request",
		"method", "POST",
	)
	return PostGroup
}

func PutGroup() slog.Attr {
	PostGroup := slog.Group(
		"request",
		"method", "PUT",
	)
	return PostGroup
}

func DeleteGroup() slog.Attr {
	PostGroup := slog.Group(
		"request",
		"method", "DELETE",
	)
	return PostGroup
}
