// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"context"
)

type Querier interface {
	AcceptTopic(ctx context.Context, arg AcceptTopicParams) error
	AcceptTopicRegistration(ctx context.Context, arg AcceptTopicRegistrationParams) error
	CreateConference(ctx context.Context, arg CreateConferenceParams) (Conference, error)
	CreateDepartment(ctx context.Context, arg CreateDepartmentParams) (Department, error)
	CreateFaculty(ctx context.Context, arg CreateFacultyParams) (Faculty, error)
	CreateGroup(ctx context.Context, arg CreateGroupParams) (Group, error)
	CreateGroupSupervisor(ctx context.Context, arg CreateGroupSupervisorParams) (GroupSupervisor, error)
	CreateLibrary(ctx context.Context, arg CreateLibraryParams) (Library, error)
	CreateNotification(ctx context.Context, arg CreateNotificationParams) (Notification, error)
	CreateStudentTopic(ctx context.Context, arg CreateStudentTopicParams) (StudentTopic, error)
	CreateTopic(ctx context.Context, arg CreateTopicParams) (Topic, error)
	CreateTopicRegistration(ctx context.Context, arg CreateTopicRegistrationParams) (TopicRegistration, error)
	CreateTopicResult(ctx context.Context, arg CreateTopicResultParams) (TopicResult, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateUserInfo(ctx context.Context, arg CreateUserInfoParams) (UserInfo, error)
	DeleteConference(ctx context.Context, id int32) error
	DeleteDepartment(ctx context.Context, id int32) error
	DeleteFaculty(ctx context.Context, id int32) error
	DeleteGroup(ctx context.Context, id int32) error
	DeleteGroupSupervisor(ctx context.Context, id int32) error
	DeleteLibrary(ctx context.Context, id int32) error
	DeleteNotification(ctx context.Context, id int32) error
	DeleteStudentTopic(ctx context.Context, id int32) error
	DeleteTopic(ctx context.Context, id int32) error
	DeleteTopicRegistration(ctx context.Context, id int32) error
	DeleteTopicResult(ctx context.Context, id int32) error
	DeleteUser(ctx context.Context, id int32) error
	DeleteUserInfo(ctx context.Context, id int32) error
	GetConference(ctx context.Context, id int32) (Conference, error)
	GetDepartment(ctx context.Context, id int32) (Department, error)
	GetFaculty(ctx context.Context, id int32) (Faculty, error)
	GetGroup(ctx context.Context, id int32) (Group, error)
	GetGroupSupervisor(ctx context.Context, id int32) (GroupSupervisor, error)
	GetLibrary(ctx context.Context, id int32) (Library, error)
	GetNotification(ctx context.Context, id int32) (Notification, error)
	GetStudentByName(ctx context.Context, name string) ([]GetStudentByNameRow, error)
	GetStudentTopic(ctx context.Context, id int32) (StudentTopic, error)
	GetTopic(ctx context.Context, id int32) (Topic, error)
	GetTopicRegistration(ctx context.Context, id int32) (TopicRegistration, error)
	GetTopicResult(ctx context.Context, id int32) (TopicResult, error)
	GetTypeAccount(ctx context.Context, id int32) (int32, error)
	GetUser(ctx context.Context, id int32) (User, error)
	GetUserByName(ctx context.Context, name string) (User, error)
	GetUserInfo(ctx context.Context, userID int32) (UserInfo, error)
	GetUserNameByID(ctx context.Context, id int32) (string, error)
	ListConferences(ctx context.Context) ([]Conference, error)
	ListDepartments(ctx context.Context) ([]Department, error)
	ListFaculties(ctx context.Context) ([]Faculty, error)
	ListGroupSupervisors(ctx context.Context) ([]GroupSupervisor, error)
	ListGroups(ctx context.Context) ([]Group, error)
	ListLibrarys(ctx context.Context) ([]Library, error)
	ListNotifications(ctx context.Context) ([]Notification, error)
	ListStudentByTopicId(ctx context.Context, topicID int32) ([]int32, error)
	ListStudentTopics(ctx context.Context) ([]StudentTopic, error)
	ListTopicRegistrations(ctx context.Context, name string) ([]TopicRegistration, error)
	ListTopicResults(ctx context.Context) ([]TopicResult, error)
	ListTopics(ctx context.Context) ([]Topic, error)
	ListTopicsFilter(ctx context.Context, name string) ([]Topic, error)
	ListUserInfos(ctx context.Context) ([]UserInfo, error)
	ListUserInfosByName(ctx context.Context, name string) ([]UserInfo, error)
	ListUsers(ctx context.Context) ([]User, error)
	UpdateConference(ctx context.Context, arg UpdateConferenceParams) error
	UpdateDepartment(ctx context.Context, arg UpdateDepartmentParams) error
	UpdateFaculty(ctx context.Context, arg UpdateFacultyParams) error
	UpdateGroup(ctx context.Context, arg UpdateGroupParams) error
	UpdateGroupSupervisor(ctx context.Context, arg UpdateGroupSupervisorParams) error
	UpdateGroupTopic(ctx context.Context, arg UpdateGroupTopicParams) error
	UpdateLibrary(ctx context.Context, arg UpdateLibraryParams) error
	UpdateNotification(ctx context.Context, arg UpdateNotificationParams) error
	UpdateStatusTopic(ctx context.Context, arg UpdateStatusTopicParams) error
	UpdateStudentTopic(ctx context.Context, arg UpdateStudentTopicParams) error
	UpdateTopicRegistration(ctx context.Context, arg UpdateTopicRegistrationParams) error
	UpdateTopicResult(ctx context.Context, arg UpdateTopicResultParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
	UpdateUserInfo(ctx context.Context, arg UpdateUserInfoParams) error
}

var _ Querier = (*Queries)(nil)
