# ⚡ Fase 5: Concorrência com Goroutines

## 💡 Conceito em 1 Minuto

**Analogia:** Restaurante

**Sem Concorrência (Sequencial):**
```text
Chef faz:
1. Prato A → 10min
2. Prato B → 10min
3. Prato C → 10min
Total: 30 minutos
```

**Com Concorrência (Goroutines):**
```text
3 Chefs trabalham ao MESMO TEMPO:
Chef 1: Prato A → 10min ┐
Chef 2: Prato B → 10min ├─ em paralelo
Chef 3: Prato C → 10min ┘
Total: 10 minutos
```

**Goroutine** = "Chef virtual" que trabalha em paralelo.

## 🎯 Sintaxe Básica

**Criar Goroutine** = Adiciona go antes da função

```go
// Sequencial (normal)
processPlayer(player1)  // espera terminar
processPlayer(player2)  // depois executa
processPlayer(player3)  // depois executa

// Concorrente (goroutines)
go processPlayer(player1)  // executa em paralelo
go processPlayer(player2)  // executa em paralelo
go processPlayer(player3)  // executa em paralelo
```

**PROBLEMA:** Como saber quando terminaram?

## 🔑 3 Patterns Essenciais

### Pattern 1: WaitGroups (Esperar que terminem)

```go
import "sync"

func main() {
    var wg sync.WaitGroup  // contador de goroutines

    players := []Player{player1, player2, player3}

    for _, p := range players {
        wg.Add(1)  // "vou lançar +1 goroutine"

        go func(player Player) {
            defer wg.Done()  // "terminei!" (quando função acaba)

            processPlayer(player)
        }(p)  // ← IMPORTANTE: passa p como argumento
    }

    wg.Wait()  // espera TODAS terminarem
    fmt.Println("Todos processados!")
}
```

Analogia:
* wg.Add(1) = "Mais 1 chef a trabalhar"
* wg.Done() = "Este chef terminou"
* wg.Wait() = "Esperar todos os chefs terminarem"

### Pattern 2: Channels (Comunicação entre goroutines)
Channel = "tubo" para passar dados entre goroutines

```go
// Criar channel
results := make(chan int)  // channel de inteiros

// Goroutine ENVIA para channel
go func() {
    result := calculateScore()
    results <- result  // ← ENVIA (bloqueia até alguém receber)
}()

// Main RECEBE do channel
score := <-results  // ← RECEBE (bloqueia até alguém enviar)
fmt.Println("Score:", score)
```

Analogia: Channel = "caixa de correio"
* `results <- 42` = "Pôr carta na caixa"
* `x := <-results` = "Tirar carta da caixa"

Exemplo Prático: Processar Stats de Múltiplos Jogadores
```go
type PlayerStats struct {
    PlayerID   int
    TotalGoals int
    Error      error
}

func calculatePlayerStats(playerID int) PlayerStats {
    // Simula cálculo demorado
    time.Sleep(1 * time.Second)

    // Calcula stats...
    return PlayerStats{
        PlayerID:   playerID,
        TotalGoals: 10,  // exemplo
    }
}

func main() {
    playerIDs := []int{1, 2, 3, 4, 5}

    // Channel para receber resultados
    results := make(chan PlayerStats, len(playerIDs))  // buffered channel

    // Lançar goroutines
    for _, id := range playerIDs {
        go func(playerID int) {
            stats := calculatePlayerStats(playerID)
            results <- stats  // envia resultado
        }(id)
    }

    // Recolher resultados
    for i := 0; i < len(playerIDs); i++ {
        stats := <-results
        fmt.Printf("Player %d: %d golos\n", stats.PlayerID, stats.TotalGoals)
    }

    // Sequencial: 5 segundos (1s × 5)
    // Concorrente: 1 segundo (todos ao mesmo tempo!)
}
```

### Pattern 3: Worker Pool (Controlar número de goroutines)

* **Problema:** Lançar 1 milhão de goroutines = crash!
* **Solução:** Limitar a N workers (ex: 10).

```go
func workerPool() {
    jobs := make(chan int, 100)     // tarefas a fazer
    results := make(chan int, 100)  // resultados

    // Criar 3 workers (3 goroutines fixas)
    numWorkers := 3
    var wg sync.WaitGroup

    for w := 1; w <= numWorkers; w++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()

            // Worker fica a buscar jobs até channel fechar
            for job := range jobs {
                fmt.Printf("Worker %d: processing job %d\n", id, job)
                time.Sleep(500 * time.Millisecond)
                results <- job * 2
            }
        }(w)
    }

    // Enviar 9 jobs
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)  // "não há mais jobs" → workers terminam

    // Esperar workers terminarem
    go func() {
        wg.Wait()
        close(results)  // fecha results quando tudo terminar
    }()

    // Recolher resultados
    for result := range results {
        fmt.Println("Result:", result)
    }
}

// Output:
// Worker 1: processing job 1
// Worker 2: processing job 2
// Worker 3: processing job 3
// Worker 1: processing job 4  ← reutiliza worker 1
// ...
```

Analogia:
* 3 chefs na cozinha (workers)
* 9 pedidos na fila (jobs)
* Cada chef pega no próximo pedido quando termina

## ⚠️ CUIDADOS Críticos

### 1. Passar variável de loop para goroutine
```go
// ❌ ERRADO
for _, player := range players {
    go func() {
        fmt.Println(player.Name)  // ← BUG! usa variável de loop
    }()
}

// ✅ CORRETO
for _, player := range players {
    go func(p Player) {  // ← passa como argumento
        fmt.Println(p.Name)
    }(player)
}
```
**Razão:** Variável `player` muda a cada iteração. Quando a goroutine executar, `player` pode já ter mudado!

### 2. Fechar channels
```go
jobs := make(chan int)

// Envia jobs
go func() {
    for i := 0; i < 10; i++ {
        jobs <- i
    }
    close(jobs)  // ✅ IMPORTANTE: fecha quando acabar
}()

// Recebe jobs
for job := range jobs {  // para quando channel fecha
    process(job)
}
```

**Regra:** Quem envia fecha o channel. Quem recebe itera até fechar.

### 3. Buffered vs Unbuffered Channels

```go
// Unbuffered (bloqueante)
ch := make(chan int)
ch <- 42  // ❌ DEADLOCK! bloqueia para sempre (ninguém a receber)

// Buffered (tem "capacidade")
ch := make(chan int, 1)  // capacidade = 1
ch <- 42  // ✅ OK! guarda no buffer
x := <-ch // recebe depois
```

Analogia:
* **Unbuffered** = Aperto de mão (sincronizado)
* **Buffered** = Caixa de correio (pode guardar N mensagens)

