# To avoid changing ps execution policy on your system, use this hackerz trickz :
# Powershell -executionpolicy Bypass .\cockroachdocker.ps1 [start|stop]

# System info
$user = $env:UserName

# Node parameters
$node1 = "roach1"
$node2 = "roach2"
$node3 = "roach3"

# Cluster parameters
$outside_port = "8080:8080"
$cluster_port = "26257:26257"
$cluster_name = "roachnet"
$database_volume = "`"//c/Users/$($user)/cockroach-data/$($node1):/cockroach/cockroach-data`""

# Docker image
$image = "cockroachdb/cockroach:v19.1.4"

function stop
{
    foreach ($node in $args) {
        Write-Host "Stopping " $node
        docker stop $node
        docker rm $node
    }
}

if ( $args[0] -eq "start" ) {
    Write-Host "Starting cluster ..."
    docker run -d --name=$node1 --hostname=$node1 --net=$cluster_name -p $cluster_port -p $outside_port -v $database_volume $image start --insecure
    if ($?) {
        docker run -d --name=$node2 --hostname=$node2 --net=$cluster_name -v $database_volume $image start start --insecure --join=$node1
        docker run -d --name=$node3 --hostname=$node3 --net=$cluster_name -v $database_volume $image start start --insecure --join=$node1
        if ($?) {
            Write-Host "Success."
        } else {
            stop $node1
        }
    }
    else {
        Write-Host "Could not start docker cluster."
    }
}

elseif ( $args[0] -eq "stop" ) {
    stop $node1 $node2 $node3
}

elseif ( $args.count -eq 0 ) {
    Write-Host "Please give an argument, either start or stop."
}

else {
    Write-Host "Argument not recognised : " $args[0]
}