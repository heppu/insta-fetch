package main

type Image struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type User struct {
	Username       string `json:"username"`
	ProfilePicture string `json:"profile_picture"`
	ID             string `json:"id"`
	FullName       string `json:"full_name"`
}

type InstaResponse struct {
	Status string `json:"status"`
	Items  []struct {
		CanDeleteComments bool   `json:"can_delete_comments"`
		Code              string `json:"code"`
		Location          struct {
			Name string `json:"name"`
		} `json:"location"`
		Images struct {
			LowResolution struct {
				Image
			} `json:"low_resolution"`
			Thumbnail struct {
				Image
			} `json:"thumbnail"`
			StandardResolution struct {
				Image
			} `json:"standard_resolution"`
		} `json:"images"`
		CanViewComments bool `json:"can_view_comments"`
		Comments        struct {
			Count int           `json:"count"`
			Data  []interface{} `json:"data"`
		} `json:"comments"`
		AltMediaURL interface{} `json:"alt_media_url"`
		Caption     struct {
			CreatedTime string `json:"created_time"`
			Text        string `json:"text"`
			From        User   `json:"from"`
			ID          string `json:"id"`
		} `json:"caption"`
		Link  string `json:"link"`
		Likes struct {
			Count int    `json:"count"`
			Data  []User `json:"data"`
		} `json:"likes"`
		CreatedTime  string `json:"created_time"`
		UserHasLiked bool   `json:"user_has_liked"`
		Type         string `json:"type"`
		ID           string `json:"id"`
		User         User   `json:"user"`
	} `json:"items"`
	MoreAvailable bool `json:"more_available"`
}
