/**
 * @Author: Joey
 * @Description:
 * @Create Date: 2023/2/21 17:39
 */

package lab

import "qf"

//
// OrderDal
//  @Description: 医嘱dal
//
type OrderDal struct {
	qf.BaseDal
}

func (dal *OrderDal) BeforeAction(kind qf.EKind, content interface{}) error {
	return nil
}

func (dal *OrderDal) AfterAction(kind qf.EKind, content interface{}) error {
	return nil
}

//
// SampleDal
//  @Description: 样本dal
//
type SampleDal struct {
	qf.BaseDal
}

func (dal *SampleDal) BeforeAction(kind qf.EKind, content interface{}) error {
	return nil
}

func (dal *SampleDal) AfterAction(kind qf.EKind, content interface{}) error {
	return nil
}

type LaboratoryDal struct {
	qf.BaseDal
}

func (dal *LaboratoryDal) BeforeAction(kind qf.EKind, content interface{}) error {
	return nil
}

func (dal *LaboratoryDal) AfterAction(kind qf.EKind, content interface{}) error {
	return nil
}

//
// CheckInDal
//  @Description: 上机dal
//
type CheckInDal struct {
	qf.BaseDal
}

func (dal *CheckInDal) BeforeAction(kind qf.EKind, content interface{}) error {
	return nil
}

func (dal *CheckInDal) AfterAction(kind qf.EKind, content interface{}) error {
	return nil
}

//
// AuditDal
//  @Description: 审核dal
//
type AuditDal struct {
	qf.BaseDal
}

func (dal *AuditDal) BeforeAction(kind qf.EKind, content interface{}) error {
	return nil
}

func (dal *AuditDal) AfterAction(kind qf.EKind, content interface{}) error {
	return nil
}

//
// ResultDal
//  @Description: 检验结果dal
//
type ResultDal struct {
	qf.BaseDal
}

func (dal *ResultDal) BeforeAction(kind qf.EKind, content interface{}) error {
	return nil
}

func (dal *ResultDal) AfterAction(kind qf.EKind, content interface{}) error {
	return nil
}

//
// GraphDal
//  @Description: 检验图像dal
//
type GraphDal struct {
	qf.BaseDal
}

func (dal *GraphDal) BeforeAction(kind qf.EKind, content interface{}) error {
	return nil
}

func (dal *GraphDal) AfterAction(kind qf.EKind, content interface{}) error {
	return nil
}

//
// ReportDal
//  @Description: 检验报告dal
//
type ReportDal struct {
	qf.BaseDal
}

func (dal *ReportDal) BeforeAction(kind qf.EKind, content interface{}) error {
	return nil
}

func (dal *ReportDal) AfterAction(kind qf.EKind, content interface{}) error {
	return nil
}
