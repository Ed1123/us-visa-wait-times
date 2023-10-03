# US Global Visa Wait Times

Scrape and get wait times for all cities in the world. It adds country to the data from official sources (https://travel.state.gov/content/travel/en/us-visas/visa-information-resources/global-visa-wait-times.html)

## Usage

```console
git clone https://github.com/Ed1123/us-visa-wait-times.git
go run .
```

- API: https://localhost:8080/wait-times
- Usable table with sorting, filtering, and searching: https://localhost:8080/table-test
- Test table created with templ: https://localhost:8080/table

## TODO

- [ ] Add countries to the data.
- [ ] Add flight prices to the data.
  - [ ] Need departure city.
- [ ] Add field for knowing if you need a visa or not.
  - [ ] Need passport country.
- [ ] Add sorting to table.
- [ ] Add filtering to table.
- [ ] Add search to table.
- [ ] Add pagination to table.
