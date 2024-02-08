# structured-logger

정형화 된 로그

- 타입이 다를 경우 생길 수 있는 데이터의 정합성 차이
- 일관된 로깅을 통해, 빠른 stack trace
- context를 통한 로거 디펜던시 제거
- 중복되어 처리되는 공통 필드(RequestID)에 대하여 context 내의 공통 필드 유지
- 로그에 대한 처리 (마스킹) 등은 struct을 커스터 마이징하여 일관되게 처리
