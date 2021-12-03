import { gql, GraphQLClient } from "graphql-request";

import { Component, h } from "preact";
import { Todo } from "./graphql";
import styles from "./TodoList.scss";

type TodoListState = {
    todos: Todo[]
}

export class TodoList extends Component<any, TodoListState> {
    constructor() {
        super()
        this.state = {
            todos: []
        }
    }

    componentWillMount() {
        this.fetchTodoList()
    }

    async fetchTodoList() {
        let client = new GraphQLClient("/query")

        // Install the VSCode GraphQL extension from the GraphQL foundation to get 
        // inline query syntax linting

        const query = gql`
        query {
            todos {
                done
                text
            }
        }
    `

        let result = await client.request(query)
        
        this.setState({
            todos: result.todos as Todo[]
        })
    }

    render() {

        let todos = this.state.todos.map(t => <div style={styles.default}>{t.text}</div>)
        return <div style={styles.container}>
            {todos}
        </div>
    }
}