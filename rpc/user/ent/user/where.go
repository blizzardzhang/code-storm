// Code generated by ent, DO NOT EDIT.

package user

import (
	"code-storm/rpc/user/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// Account applies equality check predicate on the "account" field. It's identical to AccountEQ.
func Account(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAccount, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// RealName applies equality check predicate on the "real_name" field. It's identical to RealNameEQ.
func RealName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldRealName, v))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPhone, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAddress, v))
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreateAt, v))
}

// AccountEQ applies the EQ predicate on the "account" field.
func AccountEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAccount, v))
}

// AccountNEQ applies the NEQ predicate on the "account" field.
func AccountNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldAccount, v))
}

// AccountIn applies the In predicate on the "account" field.
func AccountIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldAccount, vs...))
}

// AccountNotIn applies the NotIn predicate on the "account" field.
func AccountNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldAccount, vs...))
}

// AccountGT applies the GT predicate on the "account" field.
func AccountGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldAccount, v))
}

// AccountGTE applies the GTE predicate on the "account" field.
func AccountGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldAccount, v))
}

// AccountLT applies the LT predicate on the "account" field.
func AccountLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldAccount, v))
}

// AccountLTE applies the LTE predicate on the "account" field.
func AccountLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldAccount, v))
}

// AccountContains applies the Contains predicate on the "account" field.
func AccountContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldAccount, v))
}

// AccountHasPrefix applies the HasPrefix predicate on the "account" field.
func AccountHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldAccount, v))
}

// AccountHasSuffix applies the HasSuffix predicate on the "account" field.
func AccountHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldAccount, v))
}

// AccountEqualFold applies the EqualFold predicate on the "account" field.
func AccountEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldAccount, v))
}

// AccountContainsFold applies the ContainsFold predicate on the "account" field.
func AccountContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldAccount, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPassword, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldName, v))
}

// RealNameEQ applies the EQ predicate on the "real_name" field.
func RealNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldRealName, v))
}

// RealNameNEQ applies the NEQ predicate on the "real_name" field.
func RealNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldRealName, v))
}

// RealNameIn applies the In predicate on the "real_name" field.
func RealNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldRealName, vs...))
}

// RealNameNotIn applies the NotIn predicate on the "real_name" field.
func RealNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldRealName, vs...))
}

// RealNameGT applies the GT predicate on the "real_name" field.
func RealNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldRealName, v))
}

// RealNameGTE applies the GTE predicate on the "real_name" field.
func RealNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldRealName, v))
}

// RealNameLT applies the LT predicate on the "real_name" field.
func RealNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldRealName, v))
}

// RealNameLTE applies the LTE predicate on the "real_name" field.
func RealNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldRealName, v))
}

// RealNameContains applies the Contains predicate on the "real_name" field.
func RealNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldRealName, v))
}

// RealNameHasPrefix applies the HasPrefix predicate on the "real_name" field.
func RealNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldRealName, v))
}

// RealNameHasSuffix applies the HasSuffix predicate on the "real_name" field.
func RealNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldRealName, v))
}

// RealNameEqualFold applies the EqualFold predicate on the "real_name" field.
func RealNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldRealName, v))
}

// RealNameContainsFold applies the ContainsFold predicate on the "real_name" field.
func RealNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldRealName, v))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPhone, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldAddress, v))
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreateAt, v))
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreateAt, v))
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreateAt, vs...))
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreateAt, vs...))
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreateAt, v))
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreateAt, v))
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreateAt, v))
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreateAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}