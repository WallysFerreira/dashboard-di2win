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

window.onload = () => getCount("created_at::date", 0, null, null, "2023-10-07")

export async function getCount(group_by: string, user_id: number, tipo_documento: string | null, data_comeco: string | null, data_final: string | null) {
  const filtro = `${tipo_documento ? ", tipo_documento: \"" + tipo_documento + "\"" : ""} ${data_comeco ? ", data_comeco: \"" + data_comeco + "\"" : ""} ${data_final ? ", data_final: \"" + data_final + "\"" : ""}`

  const result = await fetch("http://localhost:8080/query", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      query: `{
        count(group_by: "${group_by}", user_id: ${user_id} ${filtro}) {
          name
          value
        }
      }`
    })
  }).then(res => res.json())

  return result
}