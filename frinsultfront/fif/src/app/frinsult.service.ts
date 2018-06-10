import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, of } from 'rxjs';
import { Insult } from './insult';
import { environment } from '../environments/environment.prod';

@Injectable({
  providedIn: 'root'
})
export class FrinsultService {
  private url = environment.insultUrl;
  
  constructor(private http: HttpClient) { }
  
  getInsults(): Observable<Insult[]> {
    return this.http.get<Insult[]>(this.url)
  }

  upvote(id: number) {
    this.http.post(this.url+"/upvote/"+id, null)
  }

  downvote(id: number) {
    this.http.post(this.url+"/downvote/"+id, null)
  }
}
