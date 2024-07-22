package bsim

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/jmoiron/sqlx"
)

type OperationalTime struct {
	StartTime time.Time
	EndTime   time.Time
}

type SystemParameterRepository struct {
	db *sqlx.DB
}

type SystemParameter struct {
	Id              int64          `db:"ID"`
	VGroup          string         `db:"VGROUP"`
	Parameter       string         `db:"PARAMETER"`
	SValue          sql.NullString `db:"SVALUE"`
	Description     sql.NullString `db:"DESCRIPTION"`
	Status          int64          `db:"STATUS"`
	CreatedDate     time.Time      `db:"CREATED_DATE"`
	CreatedBy       string         `db:"CREATED_BY"`
	LastUpdatedDate sql.NullTime   `db:"LAST_UPDATED_DATE"`
	LastUpdatedBy   sql.NullString `db:"LAST_UPDATED_BY"`
	UserLevel       int64          `db:"USER_LEVEL"`
}

type SystemParameterValue struct {
	VGroup    string `db:"VGROUP"`
	Parameter string `db:"PARAMETER"`
	SValue    string `db:"SVALUE"`
}

func NewSystemParameterRepository(db *sqlx.DB) *SystemParameterRepository {
	return &SystemParameterRepository{db: db}
}

func (s *SystemParameterRepository) GetSValueByVGroupAndParam(ctx context.Context, vgroup string, param string) (value string, err error) {
	const sql = `SELECT SVALUE FROM SYSTEM_PARAMETER WHERE VGROUP = :VGROUP AND PARAMETER = :PARAM`
	args := map[string]interface{}{
		"VGROUP": vgroup,
		"PARAM":  param,
	}
	err = Get(ctx, s.db, sql, &value, args)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (s *SystemParameterRepository) GetParamsByVGroupAndParamPattern(ctx context.Context, vgroup, param string) ([]*SystemParameterValue, error) {
	params := []*SystemParameterValue{}
	const sql = `SELECT VGROUP, PARAMETER, COALESCE(SVALUE,'') SVALUE FROM SYSTEM_PARAMETER SP JOIN STATUS S ON SP.STATUS = S.ID  WHERE SP.VGROUP = :VGROUP AND SP.PARAMETER LIKE :PARAM AND S.CODE = 'active'`
	args := map[string]interface{}{
		"VGROUP": vgroup,
		"PARAM":  param,
	}

	err := Select(ctx, s.db, sql, &params, args)
	if err != nil {
		return nil, err
	}
	return params, nil
}

func (s *SystemParameterRepository) GetParamsByVGroup(ctx context.Context, vgroup string) ([]*SystemParameterValue, error) {
	params := []*SystemParameterValue{}
	const sql = `SELECT VGROUP, PARAMETER, COALESCE(SVALUE,'') SVALUE FROM SYSTEM_PARAMETER SP JOIN STATUS S ON SP.STATUS = S.ID  WHERE SP.VGROUP = :VGROUP AND S.CODE = 'active'`
	args := map[string]interface{}{
		"VGROUP": vgroup,
	}

	err := Select(ctx, s.db, sql, &params, args)
	if err != nil {
		return nil, err
	}
	return params, nil
}

func (s *SystemParameterRepository) GetParamsByVGroupAndParameterArr(ctx context.Context, vgroup string, params []string) ([]*SystemParameterValue, error) {
	paramsResult := []*SystemParameterValue{}
	args := map[string]interface{}{
		"VGROUP": vgroup,
	}
	filter := ""

	for i, p := range params {
		index := fmt.Sprintf("PARAMETER_%d", i+1)
		args[index] = p
		filter += fmt.Sprintf(":%s, ", index)
	}

	filter = strings.TrimRight(filter, " ,")

	sql := fmt.Sprintf(`SELECT VGROUP, PARAMETER, COALESCE(SVALUE,'') SVALUE FROM SYSTEM_PARAMETER SP JOIN STATUS S ON SP.STATUS = S.ID  WHERE SP.VGROUP = :VGROUP AND PARAMETER IN (%s) AND S.CODE = 'active'`, filter)

	err := Select(ctx, s.db, sql, &paramsResult, args)
	if err != nil {
		return nil, err
	}
	return paramsResult, nil
}

func (s *SystemParameterRepository) UpdateParamsByVgroupAndParameter(ctx context.Context, vgroup string, params string, value string) error {

	args := map[string]interface{}{
		"VALUE": value,
	}
	sql := fmt.Sprintf(`UPDATE  SYSTEM_PARAMETER SET SVALUE= :VALUE WHERE VGROUP = '%s' AND PARAMETER = '%s'`, vgroup, params)

	_, err := ExecWithResult(ctx, s.db, sql, args)
	if err != nil {
		return err
	}
	return nil
}

func CallSystemParam() {
	db := ConnectDB()
	sp := NewSystemParameterRepository(db)
	desDecryptor := NewDesDecryptor()
	ctx := context.Background()
	res, err := sp.GetSValueByVGroupAndParam(ctx, "DB_MIGRATION_ORACLE_TO_PSQL", "CALENDAR_EXCEPTION_READ")
	if err != nil {
		fmt.Println("error in GetSValueByVGroupAndParam ", err)
	}
	val, err := desDecryptor.Decrypt(ctx, res)
	spew.Dump(res, err, val)
	// val = "00:00||24:00"
	// resEncr, _ := desDecryptor.Encrypt(ctx, val)
	// spew.Dump(resEncr)

	// err = sp.UpdateParamsByVgroupAndParameter(ctx, "OPERATION_TIME", "RTGS_OPERATION", resEncr)
	// spew.Dump(err)
	// res, err = sp.GetSValueByVGroupAndParam(ctx, "OPERATION_TIME", "RTGS_OPERATION")
	// if err != nil {
	// 	fmt.Println("error in GetSValueByVGroupAndParam ", err)
	// }
	// val, err = desDecryptor.Decrypt(ctx, res)
	// spew.Dump("final", res, err, val)

}

// 7b9baa6e44cddfc13d9d527b9ab43df8
