sum(1, 2);

sum(3,4,5);

sum(50,90,40);

<?php

function sum () {
    $args = func_get_arg();

    return array_reduce($args, function($acc, $item) {
        return $acc + $item;

    }, 0);
    
}