package commenthandler

type CommentRequest struct {
	Message string `json:"message"`
	PhotoID int    `json:"photo_id"`
}
