Banco de dados

CREATE TABLE courses (
	id string,
	name string,
	description string,
	category_id string
);

CREATE TABLE categories (
	id string,
	name string,
	description string
);

Consultas graphql

mutation createCategory {
  createCategory(input: {name: "tech", description: "Cursos tech"})
  {
    id,
    name,
    description
  }
}

query categories {
  categories {
    name,
    id
  }
}

mutation createCourse {
  createCourse(input: {name: "Informatica basica", description: "basico", categoryId: "<ID DA CATEGORIA>"})
  {
    id,
    name,
    description
  }
}

query courses {
  courses {
    name,
    id,
    description
  }
}

query coursesWithCategory {
  courses {
    name,
    id,
    description
    category {
      id,
      name,
      description
    }
  }
}

query categoryWithCourse {
  categories {
    name,
    id,
    description
    courses {
      id,
      name
    }
  }
}

roda no localhost:8080

$  go run github.com/99designs/gqlgen init // iniciar o projeto
$  go run github.com/99designs/gqlgen generate // gerar/atualizar 