# README

Este repositório contém materiais de estudo sobre a linguagem de programação Go (Golang). O objetivo é registar o progresso da aprendizagem dos conceitos fundamentais da linguagem, as suas características e melhores práticas de uso através de módulos organizados e avaliados pelo Claude da Anthropic.

## [Fase 1: Fundações Sólidas](MOD1.md)

* Sintaxe básica, tipos, estruturas de controlo
* **Por quê:** Go tem idiomas próprios (:=, múltiplos returns, defer) que precisam ser naturais
* **Desafio final:** CLI tool simples

## Fase 2: Ponteiros e Memória

* Ponteiros, valores vs referências, quando usar cada um
* **Por quê:** Diferente de Ruby/TS, precisas gerir isto explicitamente
* **Desafio final:** Implementar estrutura de dados própria

## Fase 3: Structs, Métodos e Interfaces

* OOP em Go, composição vs herança, interfaces implícitas
* **Por quê:** O coração do design em Go - muito diferente de classes
* **Desafio final:** Sistema com múltiplos tipos interagindo

## Fase 4: Error Handling e Panic

* Padrão if err != nil, erros custom, quando panic
* **Por quê:** Sem try/catch - precisas pensar diferente
* **Desafio final:** Função robusta com múltiplos cenários de erro

## Fase 5: Concorrência - Goroutines

* Goroutines, channels, select, patterns
* **Por quê:** A killer feature do Go, mas fácil de fazer mal
* **Desafio final:** Worker pool processando tarefas

## Fase 6: Packages e Organização

* Estrutura de projeto, módulos, visibilidade, testes
* **Por quê:** Projetos reais precisam de organização adequada
* **Desafio final:** Refatorar código em packages bem estruturados

## Fase 7: Standard Library Essencial

* io, net/http, encoding/json, context, testing
* **Por quê:** Go tem stdlib rica - usa-a antes de packages externos
* **Desafio final:** API REST básica sem frameworks

## Fase 8: Projeto Integrador

* Aplicar tudo num projeto realista
* **Por quê:** Consolidar conhecimento com algo próximo do mundo real
* **Desafio final:** API completa com testes, concorrência, persistência