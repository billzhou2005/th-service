#! /bin/bash

## 自动修改包含版本信息源码中的版本,并提交git仓库,生成版本标签,以及下一个快照版本号
## 参照maven对版本的定义,后缀为-SNAPSHOT的为开发阶段的不稳定版本
## 版本号格式为 MAJOR.MINOR.PATCH[.DESC][-SNAPSHOT],参见脚本中　RLV_FMT_REG　正则表达式定义
## 运行前要确保所有的修改都已经提交

# 获取代码中的变量定义
function get_var(){
    echo $(sed -nr "s/^\s*\/\/!\s*$1\s*(\S*).*$/\1/p" $2)    
}

# 根据代码中的变量定义的正则表达式获取代码中的变量定义的值
function get_value_by_reg(){
    echo $(sed -nr "s/^.*$1.*/\2/p" $2)    
}

# 检查版本号格式是否正确
function check_version_format(){
    if [[ -n "$1" && -z $(echo $1 | grep -E "$RLV_FMT_REG") ]] ;
    then
        echo "ERROR:invalid version string format(版本号格式错误),example: 1.0.0 OR 1.0.0.beta"
        exit -1
    fi
}

sh_folder=$(dirname $(readlink -f $0))

# 包含版本信息的源码文件，根据项目实际情况修改之
version_src=$sh_folder/verconfig/version.h

# git分支名
branch=`git branch | grep "*"`
branch=${branch/* /}
# 当前提交id
commit_id=`git rev-parse --short HEAD`

echo branch=$branch
echo commit_id=$commit_id

# 版本号格式定义
# release版本号格式
RLV_FMT_REG="([0-9]+)\.([0-9]+)\.([0-9]+)(\.([[:alnum:]]+))?"
# 开发快照版本号格式(多一个-SNAPSHOT后缀)
SS_VFMT_REG=$RLV_FMT_REG"(-(SNAPSHOT))?"

TAG_PREFIX=$(get_var @TAG_PREFIX@  $version_src)
echo TAG_PREFIX=$TAG_PREFIX
VERSION_REG=$(get_var @VERSION_REG@  $version_src)
echo VERSION_REG=$VERSION_REG
VERSION_MAJOR_REG=$(get_var @VERSION_MAJOR_REG@  $version_src)
echo VERSION_MAJOR_REG=$VERSION_MAJOR_REG
VERSION_MINOR_REG=$(get_var @VERSION_MINOR_REG@  $version_src)
echo VERSION_MINOR_REG=$VERSION_MINOR_REG
VERSION_PATCH_REG=$(get_var @VERSION_PATCH_REG@  $version_src)
echo VERSION_PATCH_REG=$VERSION_PATCH_REG
VERSION_DESC_REG=$(get_var @VERSION_DESC_REG@  $version_src)
echo VERSION_DESC_REG=$VERSION_DESC_REG
BEANCH_REG=$(get_var @BEANCH_REG@  $version_src)
echo BEANCH_REG=$BEANCH_REG
COMMIT_REG=$(get_var @COMMIT_REG@  $version_src)
echo COMMIT_REG=$COMMIT_REG

# 从源码中获取当前版本号字符串
VERSION=$(get_value_by_reg $VERSION_REG $version_src)
echo VERSION=$VERSION

if [ -z "$VERSION" ] ;
then
    echo "ERROR:failt to read version format from $version_src(从源码中读取版本信息失败)"
    exit -1
fi

# 解析版本号字符串各字段
# 主版本号
MAJOR_VERSION=$(echo $VERSION | sed -nr "s/$SS_VFMT_REG/\1/p")
# 次版本号
MINOR_VERSION=$(echo $VERSION | sed -nr "s/$SS_VFMT_REG/\2/p")
# 补丁版本号
PATCH_VERSION=$(echo $VERSION | sed -nr "s/$SS_VFMT_REG/\3/p")
# 版本号后缀
DESC=$(echo $VERSION | sed -nr "s/$SS_VFMT_REG/\5/p")
# 快照版本
SNAPSHOT=$(echo $VERSION | sed -nr "s/$SS_VFMT_REG/\7/p")

echo MAJOR_VERSION=$MAJOR_VERSION
echo MINOR_VERSION=$MINOR_VERSION
echo PATCH_VERSION=$PATCH_VERSION
echo DESC=$DESC
echo SNAPSHOT=$SNAPSHOT

[ -z "$SNAPSHOT" ] && echo "WARNING:$VERSION is not a snapshot version(当前版本号不是快照版本)"

# 版本号最末位自动加1
new_patch=$(expr $PATCH_VERSION + 1 )

echo new_patch=$new_patch

# 发行版本号(用默认值初始化)
release_version=$MAJOR_VERSION.$MINOR_VERSION.$PATCH_VERSION
# 下一个发行版本号
next_relver=$MAJOR_VERSION.$MINOR_VERSION.$new_patch
[ -n "$DESC" ] && release_version=$release_version.$DESC && next_relver=$next_relver.$DESC
#####################

# 提示用户输入发行版本号,如果输入为空则使用默认值
read -p "input release version(输入发行版本号)[$release_version]:" input_str

check_version_format $input_str

#if [[ -n "$input_str" && -z $(echo $input_str | grep -E "$RLV_FMT_REG") ]] ;
#then
#   echo "ERROR:invalid version string format(版本号格式错误),example: 1.0.0 OR 1.0.0.beta"
#   exit -1
#fi

[ -n "$input_str" ] && release_version=$input_str

echo release_version=$release_version


# 提示用户输入下一个发行版本号,如果输入为空则使用默认值
read -p "input next release version(输入下一个发行版本号)[$next_relver]:" input_str

check_version_format $input_str

#if [[ -n "$input_str" && -z $(echo $input_str | grep -E "$RLV_FMT_REG") ]] ;
#then
#   echo "ERROR:invalid version string format(版本号格式错误),example: 1.0.0 OR 1.0.0.beta"
#   exit -1
#fi

if [ "$input_str" = $release_version ] ;
then
    echo "ERROR:next version must not be same with $release_version(下一个版本号不能与上一个版本号相同)"
    exit -1
fi

[ -n "$input_str" ] && next_relver=$input_str
echo next_relver=$next_relver


# 发行版本各字段
# 主版本号
rel_major_version=$(echo $release_version | sed -nr "s/$RLV_FMT_REG/\1/p")
# 次版本号
rel_minor_version=$(echo $release_version | sed -nr "s/$RLV_FMT_REG/\2/p")
# 补丁版本号
rel_patch_version=$(echo $release_version | sed -nr "s/$RLV_FMT_REG/\3/p")
# 版本号后缀
rel_desc=$(echo $release_version | sed -nr "s/$RLV_FMT_REG/\5/p")

# 下一个快照版本各字段
# 主版本号
snap_major_version=$(echo $next_relver | sed -nr "s/$RLV_FMT_REG/\1/p")
# 次版本号
snap_minor_version=$(echo $next_relver | sed -nr "s/$RLV_FMT_REG/\2/p")
# 补丁版本号
snap_patch_version=$(echo $next_relver | sed -nr "s/$RLV_FMT_REG/\3/p")
# 版本号后缀
snap_desc=$(echo $next_relver | sed -nr "s/$RLV_FMT_REG/\5/p")
#########################
# 检查工作区是否全部有未提交的内容，有则报错退出
if [ -n "$(git status -s)" ] ;
then
    echo "ERROR:working directory not clean,can not release(工作区有未提交修改，不能执行release)"
    exit -1
fi

echo "修改源码版本号($release_version)"
sed -i -r "s!$VERSION_REG!\1$release_version\3!g" $version_src || exit -1
sed -i -r "s!$VERSION_MAJOR_REG!\1$rel_major_version\3!g" $version_src || exit -1
sed -i -r "s!$VERSION_MINOR_REG!\1$rel_minor_version\3!g" $version_src || exit -1
sed -i -r "s!$VERSION_PATCH_REG!\1$rel_patch_version\3!g" $version_src || exit -1 
sed -i -r "s!$VERSION_DESC_REG!\1$rel_desc\3!g" $version_src || exit -1
sed -i -r "s!$BEANCH_REG!\1$branch\3!g" $version_src || exit -1
sed -i -r "s!$COMMIT_REG!\1$commit_id\3!g" $version_src || exit -1

echo "提交对源码的修改 release_version=$release_version"
git add $version_src || exit -1
git commit -m "release $release_version" || exit -1

release_tag=$TAG_PREFIX$release_version
echo "发行版本标签(release_tag=$release_tag)"
git tag $release_tag || exit -1

# 下一个快照版本
next_snapshot="$next_relver-SNAPSHOT"
# 当前提交id
cur_commit_id=`git rev-parse --short HEAD`
echo "修改源码,改为快照版本号($next_snapshot),删除branch,commit_id"
sed -i -r "s!$VERSION_REG!\1$next_snapshot\3!g" $version_src || exit -1
sed -i -r "s!$VERSION_MAJOR_REG!\1$snap_major_version\3!g" $version_src || exit -1
sed -i -r "s!$VERSION_MINOR_REG!\1$snap_minor_version\3!g" $version_src || exit -1
sed -i -r "s!$VERSION_PATCH_REG!\1$snap_patch_version\3!g" $version_src || exit -1 
sed -i -r "s!$VERSION_DESC_REG!\1$snap_desc\3!g" $version_src || exit -1
sed -i -r "s!$BEANCH_REG!\1$branch\3!g" $version_src || exit -1
sed -i -r "s!$COMMIT_REG!\1$cur_commit_id\3!g" $version_src || exit -1

echo "提交对源码的修改"
git add $version_src || exit -1
git commit -m "next snapshot $next_snapshot" || exit -1

#echo "push到远端服务器"
git push || exit -1
git push origin $release_tag || exit -1