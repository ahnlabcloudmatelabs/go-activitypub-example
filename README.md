## CLI

### Create user

```bash
go run generate/main.go \
  -function=createUser \
  -id=juunini \
  -name="지상 최강의 개발자 쥬니니" \
  -bio="나는 지상 최강의 개발자 쥬니니다" \
  -icon="https://..." \
  -image="https://..."
```

### Follow

```bash
go run generate/main.go \
  -function=follow \
  -id=juunini \
  -followTarget="https://..." \
  -followTargetInboxURL="https://..."
```
