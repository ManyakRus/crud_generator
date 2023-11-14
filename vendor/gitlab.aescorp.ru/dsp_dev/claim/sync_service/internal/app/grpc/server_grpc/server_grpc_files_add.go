package server_grpc

import (
	"context"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_proto"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/files"
)

// File_FindByFileID - находит files в БД по file_id
func (s *ServerGRPC) File_FindByFileID(ctx context.Context, Request *grpc_proto.RequestString) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := files.File{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(files.File{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &files.File{}
	Model.FileID = Request.StringFind
	err = Model.Find_ByFileId()
	if err != nil {
		return &Otvet, err
	}

	// заполяем ответ
	ModelString, err := Model.GetJSON()
	if err != nil {
		return &Otvet, err
	}
	Otvet.ModelString = ModelString

	return &Otvet, err
}

// File_FindByFileID - находит files в БД по full_name
func (s *ServerGRPC) File_FindByFullName(ctx context.Context, Request *grpc_proto.RequestString) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := files.File{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(files.File{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &files.File{}
	Model.FullName = Request.StringFind
	err = Model.Find_ByFull_name()
	if err != nil {
		return &Otvet, err
	}

	// заполяем ответ
	ModelString, err := Model.GetJSON()
	if err != nil {
		return &Otvet, err
	}
	Otvet.ModelString = ModelString

	return &Otvet, err
}
