if [ $1 -eq 2 ];
then
    rm $2/$3/$4.*
elif [ $1 -eq 4 ];
then
    rm $2/$3/$4/*
    rmdir $2/$3/$4
fi