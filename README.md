# golang-rabbitmq

# we need two ports, one for admin, and one for events
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.11-management

# rabbit cli
docker exec rabbitmq rabbitmqctl

# add user
docker exec rabbitmq rabbitmqctl add_user king(username) kong(password)

# delete guest user if you want ?
docker exec rabbitmq rabbitmqctl delete_user guest

# add permissions
docker exec rabbitmq rabbitmqctl set_user_tags king administrator

# channels exchanges etc.. are resourses, they are contained in virtual hosts (kind limit and separate resourses in a logical way
# soft restriction what resourses can reach

# create virtual host ( we need to add acces to king user to vhost)
docker exec rabbitmq rabbitmqctl add_vhost customers

# set permissions (write, read, config)
docker exec rabbitmq rabbitmqctl set_permissions -p customers king ".*" ".*" ".*"