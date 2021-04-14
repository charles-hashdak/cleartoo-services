// forum-service/main.go

package main

import(
	"context"
	_ "log"
	"time"
	_ "fmt"

	pb "github.com/charles-hashdak/cleartoo-services/forum-service/proto/forum"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/mongo/options"
)

type Comment struct{
	ID 				primitive.ObjectID  `bson:"_id,omitempty"`
	UserID 			string 				`json:"user_id"`
	SubjectID 		string 				`json:"subject_id"`
	Message 		string 				`json:"message"`
	SendAt 			string 				`json:"send_at"`
	Username 		string 				`json:"username"`
	Avatar 			Photo 				`json:"avatar"`
}

type Comments []*Comment

type Subject struct{
	ID 				primitive.ObjectID  `bson:"_id,omitempty"`
	Title 			string 				`json:"title"`
	Description		string 				`json:"description"`
	Image 			Photo 				`json:"image"`
	LastComment		string 				`json:"last_comment"`
}

type Subjects []*Subject

type Photo struct{
	ID 				primitive.ObjectID 
	Url 			string
	IsMain 			bool
	Height 			int32
	Width 			int32
}

type CommentRequest struct {
	Comment 			Comment
}

type CommentResponse struct {
	Commented			bool
}

type GetCommentsRequest struct {
	SubjectID 			string
}

type GetCommentsResponse struct {
	Comments 			Comments
}

type GetSubjectsRequest struct {
}

type GetSubjectsResponse struct {
	Subjects 			Subjects
}

func MarshalCommentRequest(req *pb.CommentRequest) *CommentRequest{
	return &CommentRequest{
		Comment: 		*MarshalComment(req.Comment),
	}
}

func UnmarshalCommentRequest(req *CommentRequest) *pb.CommentRequest{
	return &pb.CommentRequest{
		Comment: 		UnmarshalComment(&req.Comment),
	}
}

func MarshalCommentResponse(req *pb.CommentResponse) *CommentResponse{
	return &CommentResponse{
		Commented: 			req.Commented,
	}
}

func UnmarshalCommentResponse(req *CommentResponse) *pb.CommentResponse{
	return &pb.CommentResponse{
		Commented: 			req.Commented,
	}
}

func MarshalGetCommentsRequest(req *pb.GetCommentsRequest) *GetCommentsRequest{
	return &GetCommentsRequest{
		SubjectID: 		req.SubjectId,
	}
}

func UnmarshalGetCommentsRequest(req *GetCommentsRequest) *pb.GetCommentsRequest{
	return &pb.GetCommentsRequest{
		SubjectId: 		req.SubjectID,
	}
}

func MarshalGetCommentsResponse(req *pb.GetCommentsResponse) *GetCommentsResponse{
	return &GetCommentsResponse{
		Comments: 			MarshalComments(req.Comments),
	}
}

func UnmarshalGetCommentsResponse(req *GetCommentsResponse) *pb.GetCommentsResponse{
	return &pb.GetCommentsResponse{
		Comments: 			UnmarshalComments(req.Comments),
	}
}

func MarshalGetSubjectsRequest(req *pb.GetSubjectsRequest) *GetSubjectsRequest{
	return &GetSubjectsRequest{
	}
}

func UnmarshalGetSubjectsRequest(req *GetSubjectsRequest) *pb.GetSubjectsRequest{
	return &pb.GetSubjectsRequest{
	}
}

func MarshalGetSubjectsResponse(req *pb.GetSubjectsResponse) *GetSubjectsResponse{
	return &GetSubjectsResponse{
		Subjects: 			MarshalSubjects(req.Subjects),
	}
}

func UnmarshalGetSubjectsResponse(req *GetSubjectsResponse) *pb.GetSubjectsResponse{
	return &pb.GetSubjectsResponse{
		Subjects: 			UnmarshalSubjects(req.Subjects),
	}
}

func MarshalComment(comment *pb.Comment) *Comment{
	objId, _ := primitive.ObjectIDFromHex(comment.Id)
	return &Comment{
		ID:				objId,
		SubjectID:		comment.SubjectId,
		UserID:			comment.UserId,
		Message:		comment.Message,
		Username:		comment.Username,
		SendAt:			comment.SendAt,
		Avatar:			*MarshalPhoto(comment.Avatar),
	}
}

func UnmarshalComment(comment *Comment) *pb.Comment{
	return &pb.Comment{
		Id:				comment.ID.Hex(),
		SubjectId:		comment.SubjectID,
		UserId:			comment.UserID,
		Message:		comment.Message,
		Username:		comment.Username,
		SendAt:			comment.SendAt,
		Avatar:			UnmarshalPhoto(&comment.Avatar),
	}
}

func MarshalComments(comments []*pb.Comment) Comments {
	collection := make(Comments, 0)
	for _, comment := range comments {
		collection = append(collection, MarshalComment(comment))
	}
	return collection
}

func UnmarshalComments(comments Comments) []*pb.Comment {
	collection := make([]*pb.Comment, 0)
	for _, comment := range comments {
		collection = append(collection, UnmarshalComment(comment))
	}
	return collection
}

func MarshalSubject(subject *pb.Subject) *Subject{
	objId, _ := primitive.ObjectIDFromHex(subject.Id)
	return &Subject{
		ID:				objId,
		Title:			subject.Title,
		Description:	subject.Description,
		Image:			*MarshalPhoto(subject.Image),
		LastComment:	subject.LastComment,
	}
}

func UnmarshalSubject(subject *Subject) *pb.Subject{
	return &pb.Subject{
		Id:				subject.ID.Hex(),
		Title:			subject.Title,
		Description:	subject.Description,
		Image:			UnmarshalPhoto(&subject.Image),
		LastComment:	subject.LastComment,
	}
}

func MarshalSubjects(subjects []*pb.Subject) Subjects {
	collection := make(Subjects, 0)
	for _, subject := range subjects {
		collection = append(collection, MarshalSubject(subject))
	}
	return collection
}

func UnmarshalSubjects(subjects Subjects) []*pb.Subject {
	collection := make([]*pb.Subject, 0)
	for _, subject := range subjects {
		collection = append(collection, UnmarshalSubject(subject))
	}
	return collection
}

func MarshalPhoto(photo *pb.Photo) *Photo{
	if(photo == nil){
		return &Photo{}
	}
	objId, _ := primitive.ObjectIDFromHex(photo.Id)
	return &Photo{
		ID:				objId,
		Url: 			photo.Url,
		IsMain:			photo.IsMain,
		Height:			photo.Height,
		Width:			photo.Width,
	}
}

func UnmarshalPhoto(photo *Photo) *pb.Photo{
	if(photo == nil){
		return &pb.Photo{}
	}
	return &pb.Photo{
		Id:				photo.ID.Hex(),
		Url: 			photo.Url,
		IsMain:			photo.IsMain,
		Height:			photo.Height,
		Width:			photo.Width,
	}
}

type repository interface{
	Comment(ctx context.Context, comment *Comment) error
	GetComments(ctx context.Context, req *GetCommentsRequest) ([]*Comment, error)
	GetSubjects(ctx context.Context, req *GetSubjectsRequest) ([]*Subject, error)
}

type MongoRepository struct{
	subjectCollection 	*mongo.Collection
	commentCollection 	*mongo.Collection
}

func (repo *MongoRepository) Comment(ctx context.Context, comment *Comment) error{
	comment.SendAt = time.Now().Format("2006-01-02 15:04:05")
	_, err := repo.commentCollection.InsertOne(ctx, comment)
	return err
}

func (repo *MongoRepository) GetComments(ctx context.Context, req *GetCommentsRequest) ([]*Comment, error){
	bsonFilters := bson.D{}
	subjectId, _ := primitive.ObjectIDFromHex(req.SubjectID)
	bsonFilters = append(bsonFilters, bson.E{"subjectid", bson.D{bson.E{"$eq", subjectId}}})
	cur, err := repo.commentCollection.Find(ctx,  bsonFilters)
	//cur, err := repo.commentCollection.Find(ctx,  bsonFilters, options.Find().SetShowRecordID(true), options.Find().SetLimit(req.Limit), options.Find().SetSkip(req.Offset))
	var comments []*Comment
	for cur.Next(ctx) {
		var comment *Comment
		if err := cur.Decode(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, err
}

func (repo *MongoRepository) GetSubjects(ctx context.Context, req *GetSubjectsRequest) ([]*Subject, error){
	bsonFilters := bson.M{}
	cur, err := repo.subjectCollection.Find(ctx, bsonFilters, nil)
	var subjects []*Subject
	for cur.Next(ctx) {
		var subject *Subject
		if err := cur.Decode(&subject); err != nil {
			return nil, err
		}
		subjects = append(subjects, subject)
	}
	return subjects, err
}