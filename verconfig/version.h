#ifndef FLCONFIG_VERSION_H_
#define FLCONFIG_VERSION_H_

// 用于　verrel.sh 的参数定义　以//!起始
//! @TAG_PREFIX@
//! @VERSION_REG@ (FL_VERSION\s+)(\S*)(.*)
//! @VERSION_MAJOR_REG@ (FL_VERSION_MAJOR\s+)(\S*)?(.*)?
//! @VERSION_MINOR_REG@ (FL_VERSION_MINOR\s+)(\S*)?(.*)?
//! @VERSION_PATCH_REG@ (FL_VERSION_PATCH\s+)(\S*)?(.*)?
//! @VERSION_DESC_REG@ (FL_VERSION_DESC\s+)(\S*)?(\s*,.*)
//! @BEANCH_REG@ (FL_SCM_BRANCH\s+)(\S*)?(\s*,.*)
//! @COMMIT_REG@ (FL_SCM_COMMIT\s+)(\S*)?(\s*,.*)

// 版本号
#define FL_VERSION 0.0.5-SNAPSHOT

#define FL_VERSION_MAJOR 0
#define FL_VERSION_MINOR 0
#define FL_VERSION_PATCH 5
#define FL_VERSION_DESC ,

// 分支名
#define FL_SCM_BRANCH main,
// commit_id
#define FL_SCM_COMMIT 5bf7be2,

#endif /* FLCONFIG_VERSION_H_ */
