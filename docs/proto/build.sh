#!bin/sh
for file in ./*.proto
do
    if test -f $file
    then
        new_dir=${file%.*}
        d1=`basename ${new_dir}`
        dest=../../protocol/${d1}
        mkdir $dest
        #echo ${dest}
        protoc -I=./ --go_out=../../ $file
    fi
    if test -d $file
    then
        echo $file "is dir"
    fi
done