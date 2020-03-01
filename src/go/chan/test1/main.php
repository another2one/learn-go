<?php

    $numCount = 20000;

    $start = microtime(true);

    function getAddNum($i){
        $sum = 0;
        while($i > 0){
            $sum+=$i;
            $i--;
        }
        return $sum;
    }

    $a = [];
    for($i = 1; $i <= $numCount; $i++){
        getAddNum($i);
    }

    $end = microtime(true);

    echo "used ". ($end - $start) . "s";

    ?>