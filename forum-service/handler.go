// forum-service/main.go

package main

import(
	"context"
	"fmt"

	pb "github.com/charles-hashdak/cleartoo-services/forum-service/proto/forum"
	userPb "github.com/charles-hashdak/cleartoo-services/user-service/proto/user"
)

type handler struct{
	repository
	userClient userPb.UserService
}

func (s *handler) SendComment(ctx context.Context, comment *pb.Comment, res *pb.CommentResponse) error {

	err := s.repository.Comment(ctx, MarshalComment(comment))

	if err != nil{
		return nil
	}

	res.Commented = true

	return nil
}

func (s *handler) GetComments(ctx context.Context, req *pb.GetCommentsRequest, res *pb.GetCommentsResponse) error {
	comments, err := s.repository.GetComment(ctx, MarshalGetCommentsRequest(req))
	if err != nil {
		return err
	}
	for _, comment := range comments {
		userRes, err2 := s.userClient.Get(ctx, &userPb.User{
			Id: comment.UserID,
		})
		if err2 != nil {
			fmt.Println(err2)
			return err2
		}
		comment.Avatar.Url = userRes.User.AvatarUrl
		comment.Username = userRes.User.Username
	}
	res.Comments = UnmarshalComments(comments)
	return nil
}

func (s *handler) GetSubjects(ctx context.Context, req *pb.GetSubjectsRequest, res *pb.GetSubjectsResponse) error {
	subjects, err := s.repository.GetSubjects(ctx, MarshalGetSubjectsRequest(req))
	if err != nil {
		return err
	}
	res.Subjects = UnmarshalSubjects(subjects)
	return nil
}