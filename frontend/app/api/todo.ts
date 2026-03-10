import { type TodoListResponse } from "~/types/todo";

export default {
	fetchTodo: async () => await useRequest<TodoListResponse>('/todos')
}