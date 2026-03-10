export interface ITodo
{
	completed: Boolean,
	id: Number,
	title: String,
	userId: Number
}

export type TodoListResponse = ITodo[];