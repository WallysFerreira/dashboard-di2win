import { group } from '@angular/animations';
import { Component } from '@angular/core';
import { environment } from 'src/environments/environment';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'dashboard';
}

export async function getCount(group_by: string, countRows: boolean, user_id: string, segment: string | null, tipo_documento: string | null, data_comeco: string | null, data_final: string | null) {
  const filtro = `${segment ? ", segmento: \"" + segment + "\"" : ""} ${tipo_documento ? ", tipo_documento: \"" + tipo_documento + "\"" : ""} ${data_comeco ? ", data_comeco: \"" + data_comeco + "\"" : ""} ${data_final ? ", data_final: \"" + data_final + "\"" : ""}`

  const result = await fetch(environment.API_URL, {
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
  const result = await fetch(environment.API_URL, {
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
  const result = await fetch(environment.API_URL, {
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

export function ajeitarNome(nome: any) {
  let nomeAjeitado = nome.split("_").map((palavra: any, index: number) => {
    palavra = palavra.toLowerCase()
      
    if (palavra.length <= 3) palavra = palavra.toUpperCase()

    if (index == 0) {
      let array = palavra.split("")
      array[0] = array[0].toUpperCase()
      palavra = array.join("")
    }

    return palavra
  }).join(" ")

  return nomeAjeitado
}
