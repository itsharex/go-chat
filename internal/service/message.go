package service

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/sirupsen/logrus"
	"go-chat/api/pb/message/v1"
	"go-chat/internal/entity"
	"go-chat/internal/logic"
	"go-chat/internal/pkg/encrypt"
	"go-chat/internal/pkg/filesystem"
	"go-chat/internal/pkg/jsonutil"
	"go-chat/internal/pkg/strutil"
	"go-chat/internal/pkg/timeutil"
	"go-chat/internal/repository/dao"
	"go-chat/internal/repository/model"
	"gorm.io/gorm"
)

type MessageService struct {
	*BaseService
	forward        *logic.MessageForwardLogic
	publisher      *logic.Publisher
	groupMemberDao *dao.GroupMemberDao
	splitUploadDao *dao.SplitUploadDao
	fileSystem     *filesystem.Filesystem
}

// SendText 文本消息
func (m *MessageService) SendText(ctx context.Context, uid int, req *message.TextMessageRequest) error {

	data := &model.TalkRecords{
		TalkType:   int(req.Receiver.TalkType),
		MsgType:    entity.MsgTypeText,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
		Content:    req.Content,
	}

	if err := m.db.Create(data).Error; err != nil {
		return err
	}

	// todo 推送MQ

	return nil
}

// SendImage 图片文件消息
func (m *MessageService) SendImage(ctx context.Context, uid int, req *message.ImageMessageRequest) error {

	parse, err := url.Parse(req.Url)
	if err != nil {
		return err
	}

	record := &model.TalkRecords{
		TalkType:   int(req.Receiver.TalkType),
		MsgType:    entity.MsgTypeFile,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
	}

	err = m.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(record).Error; err != nil {
			return err
		}

		file := &model.TalkRecordsFile{
			RecordId:     record.Id,
			UserId:       uid,
			Source:       1,
			Type:         entity.MediaFileImage,
			Drive:        entity.FileDriveMode("local"),
			OriginalName: "图片名称",
			Suffix:       strutil.FileSuffix(req.Url),
			Size:         int(req.Size),
			Path:         parse.Path,
			Url:          req.Url,
		}

		return tx.Create(file).Error
	})

	if err != nil {
		return err
	}

	return nil
}

// SendVoice 语音文件消息
func (m *MessageService) SendVoice(ctx context.Context, uid int, req *message.VoiceMessageRequest) error {

	parse, err := url.Parse(req.Url)
	if err != nil {
		return err
	}

	record := &model.TalkRecords{
		TalkType:   int(req.Receiver.TalkType),
		MsgType:    entity.MsgTypeFile,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
	}

	err = m.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(record).Error; err != nil {
			return err
		}

		file := &model.TalkRecordsFile{
			RecordId:     record.Id,
			UserId:       uid,
			Source:       1,
			Type:         entity.MediaFileAudio,
			Drive:        entity.FileDriveMode("local"),
			OriginalName: "语音文件",
			Suffix:       strutil.FileSuffix(req.Url),
			Size:         int(req.Size),
			Path:         parse.Path,
			Url:          req.Url,
		}

		return tx.Create(file).Error
	})

	if err != nil {
		return err
	}

	return nil
}

// SendVideo 视频文件消息
func (m *MessageService) SendVideo(ctx context.Context, uid int, req *message.VideoMessageRequest) error {

	parse, err := url.Parse(req.Url)
	if err != nil {
		return err
	}

	record := &model.TalkRecords{
		TalkType:   int(req.Receiver.TalkType),
		MsgType:    entity.MsgTypeFile,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
	}

	err = m.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(record).Error; err != nil {
			return err
		}

		file := &model.TalkRecordsFile{
			RecordId:     record.Id,
			UserId:       uid,
			Source:       1,
			Type:         entity.MediaFileVideo,
			Drive:        entity.FileDriveMode("local"),
			OriginalName: "语音文件",
			Suffix:       strutil.FileSuffix(req.Url),
			Size:         int(req.Size),
			Path:         parse.Path,
			Url:          req.Url,
		}

		return tx.Create(file).Error
	})

	if err != nil {
		return err
	}

	return nil
}

// SendFile 文件消息
func (m *MessageService) SendFile(ctx context.Context, uid int, req *message.FileMessageRequest) error {

	file, err := m.splitUploadDao.GetFile(uid, req.UploadId)
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("private/files/talks/%s/%s.%s", timeutil.DateNumber(), encrypt.Md5(strutil.Random(16)), file.FileExt)
	if err := m.fileSystem.Default.Copy(file.Path, filePath); err != nil {
		logrus.Error("文件拷贝失败 err: ", err.Error())
		return err
	}

	record := &model.TalkRecords{
		TalkType:   int(req.Receiver.TalkType),
		MsgType:    entity.MsgTypeFile,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
	}

	err = m.db.Transaction(func(tx *gorm.DB) error {

		if err = tx.Create(record).Error; err != nil {
			return err
		}

		data := &model.TalkRecordsFile{
			RecordId:     record.Id,
			UserId:       uid,
			Source:       1,
			Type:         entity.MediaFileOther,
			Drive:        file.Drive,
			OriginalName: file.OriginalName,
			Suffix:       file.FileExt,
			Size:         int(file.FileSize),
			Path:         filePath,
		}

		return tx.Create(data).Error
	})

	if err != nil {
		return err
	}

	return nil
}

// SendCode 代码消息
func (m *MessageService) SendCode(ctx context.Context, uid int, req *message.CodeMessageRequest) error {

	record := &model.TalkRecords{
		TalkType:   int(req.Receiver.TalkType),
		MsgType:    entity.MsgTypeCode,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
	}

	err := m.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(record).Error; err != nil {
			return err
		}

		data := &model.TalkRecordsCode{
			RecordId: record.Id,
			UserId:   uid,
			Lang:     req.Lang,
			Code:     req.Code,
		}

		return tx.Create(data).Error
	})

	if err != nil {
		return err
	}

	// todo 发送消息

	return nil
}

// SendVote 投票消息
func (m *MessageService) SendVote(ctx context.Context, uid int, req *message.VoteMessageRequest) error {
	record := &model.TalkRecords{
		TalkType:   entity.ChatGroupMode,
		MsgType:    entity.MsgTypeVote,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
	}

	options := make(map[string]string)
	for i, value := range req.Options {
		options[fmt.Sprintf("%c", 65+i)] = value
	}

	num := m.groupMemberDao.CountMemberTotal(int(req.Receiver.ReceiverId))

	err := m.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(record).Error; err != nil {
			return err
		}

		data := &model.TalkRecordsVote{
			RecordId:     record.Id,
			UserId:       uid,
			Title:        req.Title,
			AnswerMode:   int(req.Mode),
			AnswerOption: jsonutil.Encode(options),
			AnswerNum:    int(num),
			IsAnonymous:  int(req.Anonymous),
		}

		return tx.Create(data).Error
	})

	if err != nil {
		return err
	}

	// todo 发消息

	return nil
}

// SendEmoticon 表情消息
func (m *MessageService) SendEmoticon(ctx context.Context, uid int, req *message.EmoticonMessageRequest) error {

	emoticon := &model.EmoticonItem{}
	if err := m.db.Model(&model.EmoticonItem{}).Where("id = ?", req.EmoticonId).First(emoticon).Error; err != nil {
		return err
	}

	if emoticon.UserId > 0 && emoticon.UserId != uid {
		return errors.New("表情包不存在！")
	}

	record := &model.TalkRecords{
		TalkType:   int(req.Receiver.TalkType),
		MsgType:    entity.MsgTypeFile,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
	}

	err := m.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(record).Error; err != nil {
			return err
		}

		if err := tx.Create(&model.TalkRecordsFile{
			RecordId:     record.Id,
			UserId:       uid,
			Source:       2,
			Type:         entity.GetMediaType(emoticon.FileSuffix),
			OriginalName: "图片表情",
			Suffix:       emoticon.FileSuffix,
			Size:         emoticon.FileSize,
			Path:         emoticon.Url,
			Url:          emoticon.Url,
		}).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// SendForward 转发消息
func (m *MessageService) SendForward(ctx context.Context, uid int, req *message.ForwardMessageRequest) error {

	var (
		err   error
		items []*logic.ForwardRecord
	)

	// 验证转发消息合法性
	if err := m.forward.Verify(ctx, uid, req); err != nil {
		return err
	}

	if req.Mode == 1 {
		items, err = m.forward.MultiMergeForward(ctx, uid, req)
	} else {
		items, err = m.forward.MultiSplitForward(ctx, uid, req)
	}

	fmt.Println(items, err)

	return nil
}

// SendLocation 位置消息
func (m *MessageService) SendLocation(ctx context.Context, uid int, req *message.LocationMessageRequest) error {

	record := &model.TalkRecords{
		TalkType:   int(req.Receiver.TalkType),
		MsgType:    entity.MsgTypeLocation,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
	}

	err := m.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(record).Error; err != nil {
			return err
		}

		data := &model.TalkRecordsLocation{
			RecordId:  record.Id,
			UserId:    uid,
			Longitude: req.Longitude,
			Latitude:  req.Latitude,
		}

		return tx.Create(data).Error
	})

	if err != nil {
		return err
	}

	// todo 推送数据

	return nil
}

// SendBusinessCard 推送用户名片消息
func (m *MessageService) SendBusinessCard(ctx context.Context, uid int) error {
	return nil
}

// SendLogin 推送用户登录消息
func (m *MessageService) SendLogin(ctx context.Context, uid int, req *message.LoginMessageRequest) error {

	record := &model.TalkRecords{
		TalkType:   entity.ChatPrivateMode,
		MsgType:    entity.MsgTypeLogin,
		UserId:     4257, // 机器人ID
		ReceiverId: uid,
	}

	err := m.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(record).Error; err != nil {
			return err
		}

		data := &model.TalkRecordsLogin{
			RecordId: record.Id,
			UserId:   uid,
			Ip:       req.Ip,
			Platform: req.Platform,
			Agent:    req.Agent,
			Address:  req.Address,
			Reason:   req.Reason,
		}

		return tx.Create(data).Error
	})

	if err != nil {
		return err
	}

	// todo 发消息

	return nil
}

func (m *MessageService) onPushMessage(ctx context.Context) {
	_ = m.publisher.Publish(ctx, "", "")
}