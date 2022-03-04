/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mysql_dao

import (
	"context"
	"database/sql"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/media/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type VideoSizesDAO struct {
	db *sqlx.DB
}

func NewVideoSizesDAO(db *sqlx.DB) *VideoSizesDAO {
	return &VideoSizesDAO{db}
}

// Insert
// insert into video_sizes(video_size_id, size_type, volume_id, local_id, secret, width, height, file_size, video_start_ts, file_path) values (:video_size_id, :size_type, :volume_id, :local_id, :secret, :width, :height, :file_size, :video_start_ts, :file_path)
// TODO(@benqi): sqlmap
func (dao *VideoSizesDAO) Insert(ctx context.Context, do *dataobject.VideoSizesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into video_sizes(video_size_id, size_type, volume_id, local_id, secret, width, height, file_size, video_start_ts, file_path) values (:video_size_id, :size_type, :volume_id, :local_id, :secret, :width, :height, :file_size, :video_start_ts, :file_path)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

// InsertTx
// insert into video_sizes(video_size_id, size_type, volume_id, local_id, secret, width, height, file_size, video_start_ts, file_path) values (:video_size_id, :size_type, :volume_id, :local_id, :secret, :width, :height, :file_size, :video_start_ts, :file_path)
// TODO(@benqi): sqlmap
func (dao *VideoSizesDAO) InsertTx(tx *sqlx.Tx, do *dataobject.VideoSizesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into video_sizes(video_size_id, size_type, volume_id, local_id, secret, width, height, file_size, video_start_ts, file_path) values (:video_size_id, :size_type, :volume_id, :local_id, :secret, :width, :height, :file_size, :video_start_ts, :file_path)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

// SelectByFileLocation
// select id, video_size_id, size_type, volume_id, local_id, secret, width, height, file_size, video_start_ts, file_path from video_sizes where volume_id = :volume_id and local_id = :local_id
// TODO(@benqi): sqlmap
func (dao *VideoSizesDAO) SelectByFileLocation(ctx context.Context, volume_id int64, local_id int32) (rValue *dataobject.VideoSizesDO, err error) {
	var (
		query = "select id, video_size_id, size_type, volume_id, local_id, secret, width, height, file_size, video_start_ts, file_path from video_sizes where volume_id = ? and local_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, volume_id, local_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByFileLocation(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.VideoSizesDO{}
	if rows.Next() {
		// TODO(@benqi): not use reflect
		err = rows.StructScan(do)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByFileLocation(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

// SelectSecret
// select secret from video_sizes where volume_id = :volume_id and local_id = :local_id limit 1
// TODO(@benqi): sqlmap
func (dao *VideoSizesDAO) SelectSecret(ctx context.Context, volume_id int64, local_id int32) (rValue int64, err error) {
	var query = "select secret from video_sizes where volume_id = ? and local_id = ? limit 1"
	err = dao.db.Get(ctx, &rValue, query, volume_id, local_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("get in SelectSecret(_), error: %v", err)
	}

	return
}

// SelectListByVideoSizeId
// select id, video_size_id, size_type, volume_id, local_id, secret, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id = :video_size_id order by local_id asc
// TODO(@benqi): sqlmap
func (dao *VideoSizesDAO) SelectListByVideoSizeId(ctx context.Context, video_size_id int64) (rList []dataobject.VideoSizesDO, err error) {
	var (
		query = "select id, video_size_id, size_type, volume_id, local_id, secret, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id = ? order by local_id asc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, video_size_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByVideoSizeId(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.VideoSizesDO
	for rows.Next() {
		v := dataobject.VideoSizesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectListByVideoSizeId(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectListByVideoSizeIdWithCB
// select id, video_size_id, size_type, volume_id, local_id, secret, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id = :video_size_id order by local_id asc
// TODO(@benqi): sqlmap
func (dao *VideoSizesDAO) SelectListByVideoSizeIdWithCB(ctx context.Context, video_size_id int64, cb func(i int, v *dataobject.VideoSizesDO)) (rList []dataobject.VideoSizesDO, err error) {
	var (
		query = "select id, video_size_id, size_type, volume_id, local_id, secret, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id = ? order by local_id asc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, video_size_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByVideoSizeId(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.VideoSizesDO
	for rows.Next() {
		v := dataobject.VideoSizesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectListByVideoSizeId(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}