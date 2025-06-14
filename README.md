## 디버깅 실행 방법

1. VSCode에서 Dev Container로 열기 (오른쪽 하단 > Reopen in Container)
2. 왼쪽 디버그 탭 > "Debug in Dev Container" 선택
3. 브레이크포인트 설정 후 F5

## 빌드

- 리눅스 타겟 빌드: main.go가 위치한 디렉터리에서 ```go build``` 실행
- 윈도우 타겟 빌드: main.go가 위치한 디렉터리에서 ```GOOS=windows GOARCH=amd64 go build``` 실행

## 리눅스 systemd 서비스로 등록

```bash
cd /etc/systemd/system
touch monitoring.service
vi monitoring.service
```

다음 내용 입력 후 :wq
```bash
[Unit]
Description=Monitoring App
After=network.target

[Service]
ExecStart=[실행파일 경로]
Restart=always
RestartSec=3

[Install]
WantedBy=multi-user.target
```

```bash
systemctl daemon-reexec
systemctl enable monitoring
systemctl start monitoring
```

## 실시간 로그 확인
```bash
watch -n 1 "journalctl -u monitoring -n 1 -f --output=cat --no-pager"
```
