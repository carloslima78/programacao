
# Python

Este artigo é projetado para ser executado em um ambiente Ubuntu. Certifique-se de ter o Python instalado no seu sistema Ubuntu. Se não tiver, você pode instalar utilizando o comando sudo apt-get install python3.

## Módulo 1: Introdução e Configuração

Vamos começar com o básico - instalando o Python e configurando o ambiente de desenvolvimento.

1. Instalação do Python
Certifique-se de ter o Python instalado. Você pode verificar digitando python3 --version no seu terminal. Se não estiver instalado, você pode instalar com o seguinte comando:

sudo apt-get update
sudo apt-get install python3

2. Instalação do Editor de Código (Opcional)
Um bom editor de código facilita o desenvolvimento. Algumas opções populares incluem o Visual Studio Code, Sublime Text ou Atom. Você pode instalá-los usando o seguinte comando:


sudo snap install code --classic  # Visual Studio Code
sudo snap install sublime-text --classic  # Sublime Text
sudo snap install atom --classic  # Atom

## Módulo 2: Conceitos Básicos
Agora que temos o ambiente configurado, vamos aprender alguns conceitos básicos de Python.

1. Hello, World!

Vamos começar com o famoso "Hello, World!" para garantir que tudo está funcionando corretamente.

Crie um arquivo chamado hello_world.py com o seguinte conteúdo:

print("Hello, World!")

Execute o script no terminal com o seguinte comando:

python3 hello_world.py

Você deve ver a saída Hello, World! no seu terminal.

2. Variáveis e Tipos de Dados

Em Python, você não precisa declarar explicitamente o tipo de uma variável. O interpretador Python faz isso por você.

# Exemplo de variáveis
nome = "João"
idade = 25
altura = 1.75
e_estudante = True

# Imprimindo valores das variáveis
print(nome)
print(idade)
print(altura)
print(e_estudante)

3. Condicionais (if, elif, else)

# Exemplo de condicionais
nota = 85

if nota >= 90:
    print("A")
elif nota >= 80:
    print("B")
elif nota >= 70:
    print("C")
else:
    print("D")

4. Loops (for, while)

# Exemplo de loops
# Loop for
frutas = ["maçã", "banana", "laranja"]
for fruta in frutas:
    print(fruta)

# Loop while
contador = 0
while contador < 5:
    print(contador)
    contador += 1

5. Funções

# Exemplo de função
def saudacao(nome):
    print("Olá,", nome)

saudacao("Maria")

## Módulo 3: Estruturas de Dados

1. Listas

# Exemplo de listas
lista_numeros = [1, 2, 3, 4, 5]
print(lista_numeros[0])  # Imprime o primeiro elemento da lista
print(len(lista_numeros))  # Imprime o tamanho da lista

2. Dicionários

# Exemplo de dicionários
dados_pessoa = {"nome": "João", "idade": 30, "cidade": "São Paulo"}
print(dados_pessoa["idade"])  # Imprime a idade

3. Tuplas

# Exemplo de tuplas
coordenadas = (10, 20)
x, y = coordenadas
print(x)  # Imprime 10
print(y)  # Imprime 20

## Módulo 4: Bibliotecas Úteis

1. PyAutoGUI

O PyAutoGUI é uma biblioteca para automação de tarefas no computador.

Para instalar:

pip3 install pyautogui

Exemplo de uso:

import pyautogui

pyautogui.moveTo(100, 100, duration=1)  # Move o cursor para (100, 100) em 1 segundo
pyautogui.click()  # Clica no local atual do cursor

2. Requests (para requisições HTTP)

O Requests é uma biblioteca HTTP para Python.

Para instalar:

pip3 install requests

Exemplo de uso:

```python
import requests

response = requests.get('https://api.exemplo.com/dados')
print(response.text)
```
3. Boto3

Biblioteca popular do Python utilizada para interagir com serviços da Amazon Web Services (AWS). Com o boto3, você pode criar, configurar e gerenciar recursos da AWS diretamente do seu código Python. Isso inclui serviços como armazenamento em nuvem (Amazon S3), computação em nuvem (Amazon EC2), banco de dados (Amazon RDS), entre outros.

Para instalar:

```bash
pip install boto3
```

Exemplo de uso:

- Listando os buckets do Amazon S3:

```python
import boto3

# Cria uma instância do cliente S3
s3 = boto3.client('s3')

# Lista os buckets
response = s3.list_buckets()

# Itera sobre os buckets e imprime seus nomes
for bucket in response['Buckets']:
    print(bucket['Name'])
```

- Enviando um Arquivo para o Amazon S3:

```python
import boto3

# Cria uma instância do cliente S3
s3 = boto3.client('s3')

# Envia um arquivo local para o Amazon S3
bucket_name = 'nome-do-seu-bucket'
file_name = 'caminho-do-seu-arquivo-local/arquivo.txt'
object_name = 'arquivo.txt'  # Nome do objeto no S3
s3.upload_file(file_name, bucket_name, object_name)
```
- Listando os Arquivos de um Bucket no Amazon S3:

```python
import boto3

# Cria uma instância do cliente S3
s3 = boto3.client('s3')

# Envia um arquivo local para o Amazon S3
import boto3

# Cria uma instância do cliente S3
s3 = boto3.client('s3')

# Lista os objetos em um bucket
bucket_name = 'nome-do-seu-bucket'
response = s3.list_objects_v2(Bucket=bucket_name)

# Itera sobre os objetos e imprime seus nomes
for obj in response['Contents']:
    print(obj['Key'])
```

### Principais Recursos

- **Documentação Oficial**: A documentação oficial do boto3 é uma excelente fonte de referência para aprender sobre todos os serviços da AWS suportados e suas operações disponíveis. Você pode encontrá-la em [AWS SDK for Python (Boto3) documentation](https://boto3.amazonaws.com/v1/documentation/api/latest/index.html).

- **Exemplos e Tutoriais**: A comunidade da AWS tem muitos exemplos e tutoriais para ajudar a entender como usar o boto3 para diferentes casos de uso. A [AWS Code Examples Repository](https://github.com/awsdocs/aws-doc-sdk-examples) é um bom lugar para começar.

O boto3 é uma ferramenta poderosa para automatizar tarefas e interagir com os serviços da AWS de forma programática. À medida que você explora e utiliza a biblioteca, você descobrirá uma ampla gama de funcionalidades que podem ser utilizadas para atender às suas necessidades de desenvolvimento na nuvem.


