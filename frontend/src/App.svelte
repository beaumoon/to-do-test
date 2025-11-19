<script lang="ts">
  import Todo from "./lib/Todo.svelte";
  import type { TodoItem } from "./lib/types";

  let todos: TodoItem[] = $state([]);

<<<<<<< HEAD
=======
  // initiate variables to store the user's input title & description
  let new_title = "";
  let new_description = "";

>>>>>>> b6e6ab2 (Completed test)
  async function fetchTodos() {
    try {
      const response = await fetch("http://localhost:8080/");
      if (response.status !== 200) {
        console.error("Error fetching data. Response status not 200");
        return;
      }

      todos = await response.json();
    } catch (e) {
      console.error("Could not connect to server. Ensure it is running.", e);
    }
  }

<<<<<<< HEAD
=======
  async function writeTodo(title, description) {

    try {
      // send a post request to backend listener
      const response = await fetch("http://localhost:8080/", {method: "POST", headers: {"Content-Type": "application/json"}, body: JSON.stringify({ title, description })});
      
      if (response.status !== 200) {
        console.error("Error fetching data. Response status not 200");
        return;
      }
      
    } 
    
    catch (e) {
      console.error("Could not connect to server. Ensure it is running.", e);
    }

    // refresh list with new to-do added
    fetchTodos();

    // somehow reset input boxes to be empty for QoL- svelte equivalent of document.getElementById("title").value = '';

  }

>>>>>>> b6e6ab2 (Completed test)
  // Initially fetch todos on page load
  $effect(() => {
    fetchTodos();
  });
</script>

<<<<<<< HEAD
=======


>>>>>>> b6e6ab2 (Completed test)
<main class="app">
  <header class="app-header">
    <h1>TODO</h1>
  </header>

  <div class="todo-list">
    {#each todos as todo}
      <Todo title={todo.title} description={todo.description} />
    {/each}
  </div>

<<<<<<< HEAD
  <h2 class="todo-list-form-header">Add a Todo</h2>
  <form class="todo-list-form">
    <input placeholder="Title" name="title" />
    <input placeholder="Description" name="description" />
    <button>Add Todo</button>
=======
  <!-- bound title & description input to variables to use as parameters in backend-pass function -->
  <h2 class="todo-list-form-header">Add a Todo</h2>
  <form class="todo-list-form">
    <input placeholder="Title" name="title" bind:value={new_title}/>
    <input placeholder="Description" name="description" bind:value={new_description}/>
    <button type="button" on:click={() => writeTodo(new_title, new_description)}>Add Todo</button>
>>>>>>> b6e6ab2 (Completed test)
  </form>
</main>

<style>
  .app {
    color: white;
    background-color: #282c34;

    text-align: center;
    font-size: 24px;

    min-height: 100vh;
    padding: 20px;
  }

  .app-header {
    font-size: calc(10px + 4vmin);
    margin-top: 50px;
  }

  .todo-list {
    margin: 50px 100px 0px 100px;
  }

  .todo-list-form-header {
    margin-top: 100px;
  }

  .todo-list-form {
    margin-top: 10px;
  }
</style>
