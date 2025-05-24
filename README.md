## 디버깅 실행 방법

1. VSCode에서 Dev Container로 열기 (오른쪽 하단 > Reopen in Container)
2. 왼쪽 디버그 탭 > "Debug in Dev Container" 선택
3. 브레이크포인트 설정 후 F5

## 빌드

- 리눅스 타겟 빌드: main.go가 위치한 디렉터리에서 `go build` 실행
- 윈도우 타겟 빌드: main.go가 위치한 디렉터리에서 `GOOS=windows GOARCH=amd64 go build` 실행

## 리눅스 systemd 서비스로 등록

```
cd /etc/systemd/system
sudo touch monitoring.service
sudo vi monitoring.service
```

다음 내용 입력 후 :wq
```
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

```
sudo systemctl daemon-reexec
sudo systemctl enable monitoring
sudo systemctl start monitoring
```

## 실시간 로그 확인
```
journalctl -u monitoring -f
```

## 로그를 실시간으로 파일 생성
```
journalctl -u monitoring -f >> [원하는 경로 + 파일명]
```




