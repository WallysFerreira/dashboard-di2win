## Sobre as imagens docker

A imagem do **backend** precisa de 3 variaveis
- ALLOWED_ORIGIN: Endereço do front-end
- POSTGRE_URL: Endereço do banco de dados
- TEST_POSTGRE_URL (opcional): Endereço do banco de dados que vai ser usado nos testes unitários

A imagem do **frontend** precisa de 1 variavel
- API_URL: Endereço do back-end