package handlers

// TODO possibly define what types can data be

type Response struct {
	Data any     `json:"data"`
	Err  *string `json:"error"`
}

type SubmitLinkInput struct {
	// TODO validate string
	OriginalURL string `json:"original_url"`
}
