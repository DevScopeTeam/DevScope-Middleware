#!/bin/sh

# 变量
ProjectName=DevScope-Middleware
url=api.devscope.search.ren
region=cn-shanghai
dir=../Deploy/${ProjectName}

swag fmt
swag init

if [ ! -d "${dir}/code" ]; then
    mkdir ${dir}/code
fi

rm -rf ${dir}/code/*
cp -r * ${dir}/code
cd ${dir}

mv ./code/s.yaml ./
rm ./code/push.sh

# 检查系统环境是linux还是mac
if [[ "$OSTYPE" == "linux-gnu" ]]; then
    sed -i 's/{{ ProjectName }}/'${ProjectName}'/g' s.yaml
    sed -i 's/{{ region }}/'${region}'/g' s.yaml
    sed -i 's/127.0.0.1:9000/'${url}'/g' code/docs/docs.go
    sed -i 's/127.0.0.1:9000/'${url}'/g' code/docs/swagger.json
    sed -i 's/127.0.0.1:9000/'${url}'/g' code/docs/swagger.yaml
elif [[ "$OSTYPE" == "darwin"* ]]; then
    # 检查是否安装了gsed
    if ! command -v gsed &> /dev/null
    then
        echo "gsed could not be found"
        echo "please install gsed"
        exit
    fi
    gsed -i 's/{{ ProjectName }}/'${ProjectName}'/g' s.yaml
    gsed -i 's/{{ region }}/'${region}'/g' s.yaml
    gsed -i 's/127.0.0.1:9000/'${url}'/g' code/docs/docs.go
    gsed -i 's/127.0.0.1:9000/'${url}'/g' code/docs/swagger.json
    gsed -i 's/127.0.0.1:9000/'${url}'/g' code/docs/swagger.yaml
fi

git add .
git commit -m "update"
git push