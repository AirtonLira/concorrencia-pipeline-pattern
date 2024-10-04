
#  CONCORRENCIA-PIPELINE-PATTERN

Este  contém um exemplo prático do uso de  goroutines e pipelines em  Go. A  é demonstrar como usar a ⚔️ concorrência e criar pipelines eficientes para ⚙️ processar 📊, explorando o 🏋️‍♂️ poder do Go para criar 🖥️ mais 🚀 rápidos e escaláveis.

## 📚 Sobre o Projeto

O projeto consiste em uma implementação onde várias goroutines são utilizadas para converter imagens para cinza, formando um orquestrador de pipeline. Cada etapa do pipeline é realizada de forma assertiva, aproveitando ao máximo desse padrão.

A 💡 principal é mostrar como:
- Criar goroutines em 🐹 Go
- 🔄 Sincronizar usando 📨 canais
- Implementar orquestração pipelines para conversão das imagens de maneira eficiente
- Utilizar padrões idiomáticos do Go
- Carregar, ler, converter para cinza e salvar imagens JPEG 

## 🗂️ Estrutura do Projeto

- `main.go`: 📄 principal que contém a configuração do pipeline e a orquestração de inicialização das goroutines.
- `internal/images/images.go`: Implementação do pipeline de carregar, ler, converter e gerar a nova imagem, onde cada etapa é modularizada para facilitar a 🔧 manutenção.
- `internal/images/*.jpg`: Imagens de exemplo utilizadas no projeto.
- `README.md`: 📄 do projeto, instruções de instalação e  execução.

## 🛠️ Tecnologias Utilizadas

- 🖥️: 🐹 Go (Golang)
- 📁 Versionamento:  Git

## ▶️ Como Executar o Projeto

1. 🌀 Clone o projeto:
   ```sh
   git clone https://github.com/code-heim/go_21_goroutines_pipeline.git
   ```

2. Navegue até a pasta do projeto:
   ```sh
   cd concorrencia-pipeline-pattern
   ```

3. ▶️ Execute o projeto:
   ```sh
   go run app/cmd/main.go
   ```

## 🤝 Contribuindo

Fique à vontade para 🤲 contribuir com este projeto. Caso encontre 🐞 problemas ou queira sugerir 💡 melhorias, siga os passos abaixo:

1. 🍴 Fork o 📁
2. Crie uma nova  branch: `git checkout -b minha-feature`
3. Faça suas alterações e commite: `git commit -m 'Minha nova 💡 feature'`
4. Envie para o seu 🍴 fork: `git push origin minha-feature`
5. Abra um 📬 Pull Request
