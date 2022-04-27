1. 安装redis, mysql

```bash
paru -S mariadb
sudo mariadb-install-db --user=mysql --basedir=/usr --datadir=/var/lib/mysql
sudo systemctl enable mariadb
sudo systemctl start mariadb
# 配置一些安全选项
sudo mysql_secure_installation
# 添加用户
# 给用户赋予权限
sudo mysql -u root -p
MariaDB [(none)]> CREATE USER 'bobo'@'localhost' IDENTIFIED BY '8888';
MariaDB [(none)]> GRANT ALL PRIVILEGES ON *.* TO 'bobo'@'localhost';
MariaDB [(none)]> FLUSH PRIVILEGES;
MariaDB [(none)]> quit


# redis
paru -S redis
sudo systemctl enable redis
sudo systemctl start redis
```

2. 复制并更新本地配置：

```bash
cp configs/redis.toml.example configs/redis.toml && cp configs/mysql.toml.example configs/mysql.toml && cp configs/http.toml.example configs/http.toml
```

3. 启动服务

```bash
// 注意，util.env.go中定义了一些环境变量，需要根据实际情况来修改
// 下面这个定义，既是configs/source/congtu.json，又是上游仓库scale.go中network/congtu.json的值
❯ go build -o ./cmd/subscan -v ./cmd

./subscan start substrate
```
