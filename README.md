# image-service

- RESTful APIs running on `localhost:80`

---

### Test

<details>
<summary><code>GET</code> <code><b>/</b></code></summary>

##### Responses

> | http code | content-type | response                     |
> | --------- | ------------ | ---------------------------- |
> | `200`     | `text/plain` | `Hello from a HandleFunc #1` |

</details>

---

### Upload images

<details>
<summary><code>POST</code> <code><b>/upload</b></code> <code>(Upload image to image server)</code></summary>

##### Body (form-data)

> | key   | required | data type | description                       |
> | ----- | -------- | --------- | --------------------------------- |
> | image | true     | file      | The content type should be image. |

##### Params

> | key        | required | description                |
> | ---------- | -------- | -------------------------- |
> | collection | true     | The name of the collection |

##### Responses

> | http code    | content-type       | response                                                              |
> | ------------ | ------------------ | --------------------------------------------------------------------- |
> | `200`        | `application/json` | `{"message": "Image uploaded successfully", "id": "id of the image"}` |
> | `405`, `500` | `application/json` | `{"message": "Failed", "error": "error message"}`                     |

</details>

---

### Display images

<details>
<summary><code>GET</code> <code><b>/upload</b></code> <code>(Upload image to image server)</code></summary>

##### Params

> | key        | required | description                |
> | ---------- | -------- | -------------------------- |
> | \_id       | true     | The id of the image        |
> | collection | true     | The name of the collection |

##### Responses

> | http code    | content-type       |                                                   |
> | ------------ | ------------------ | ------------------------------------------------- |
> | `200`        | `image/*`          |                                                   |
> | `405`, `500` | `application/json` | `{"message": "Failed", "error": "error message"}` |

</details>

---

### Delete images

<details>
<summary><code>DELETE</code> <code><b>/delete</b></code> <code>(Delete a specific image)</code></summary>

##### Params

> | key        | required | description                |
> | ---------- | -------- | -------------------------- |
> | \_id       | true     | The id of the image        |
> | collection | true     | The name of the collection |

##### Responses

> | http code    | content-type       | response                                                               |
> | ------------ | ------------------ | ---------------------------------------------------------------------- |
> | `200`        | `application/json` | `{"message": "Deleted %d documents", "id": "Id of deleted documents"}` |
> | `405`, `500` | `application/json` | `{"message": "Failed", "error": "error message"}`                      |

</details>
