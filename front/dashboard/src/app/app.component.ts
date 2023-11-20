import { group } from '@angular/animations';
import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'dashboard';
}

export async function getCount(group_by: string, countRows: boolean, user_id: string, tipo_documento: string | null, data_comeco: string | null, data_final: string | null) {
  const filtro = `${tipo_documento ? ", tipo_documento: \"" + tipo_documento + "\"" : ""} ${data_comeco ? ", data_comeco: \"" + data_comeco + "\"" : ""} ${data_final ? ", data_final: \"" + data_final + "\"" : ""}`

  const result = await fetch("http://localhost:8080/query", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      query: `{
        count(group_by: "${group_by}", count_rows: ${countRows}, user_id: ${user_id} ${filtro}) {
          name
          value
        }
      }`
    })
  }).then(res => res.json())

  return result
}

export async function getUsers() {
  const result = await fetch("http://localhost:8080/query", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      query: `{
        user {
          id
          name
          segment
        }
      }`
    })
  }).then(res => res.json())

  return result
}

export async function getExtracts() {
  const result = await fetch("http://localhost:8080/query", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      query: `{
        extract {
          id
          created_at
          pages_processed
          doc_type
          user_id
        }
      }`
    })
  }).then(res => res.json())

  return result
}
