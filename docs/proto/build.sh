#!bin/sh
for file in ./*.proto
do
    if test -f $file
    then
        echo $file
        new_dir=${file%.*}
        d1=`basename ${new_dir}`
        dest=../../protocol/${d1}
        new_package=`basename $file`
        echo $new_package
        new_package=${new_package%.*}
        echo $new_package
        package=${new_package//_/}
        #echo ${package,,}
        package=`echo ${package} | tr 'A-Z' 'a-z'`
        echo $package
        #mkdir $dest
        #echo ${dest}
        sed -i -e "s/github.com\/futuopen\/ftapi4go\/pb\/${package}/github.com\/mymmsc\/futu4g\/protocol\/${new_package}/g" $file
        rm ${file}-e
        protoc -I=./ --go_out=../../ $file
    fi
    if test -d $file
    then
        echo $file "is dir"
    fi
done