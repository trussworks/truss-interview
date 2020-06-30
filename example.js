const rowHeaders = [
  'Name',
  'Climate',
  'Residents',
  'Terrain',
  'Population',
  'Surface Water',
]

const rowDataFns = [
  (r) => renderText(r.name),
  (r) => renderText(r.climate),
  (r) => renderText(r.residents.length),
  (r) => renderList(r.terrain),
  (r) => renderText(numeral(r.population).format('0, 0')),
  (r) => renderText(calcSurfaceWater(r)),
]

function start() {
  const apiUrl = 'https://swapi.dev/api/planets/'

  fetch(apiUrl)
    .then((response) => response.json())
    .then((data) => renderTable(data))
}

function renderList(itemsString) {
  const itemsArr = itemsString.split(', ')
  const list = document.createElement('ul')
  itemsArr.forEach((c) => {
    const item = document.createElement('li')
    item.appendChild(document.createTextNode(c))
    list.appendChild(item)
  })

  return list
}

function renderText(string) {
  let display = string
  if (string === 'unknown') display = '?'
  return document.createTextNode(display)
}

function calcSurfaceWater(planet) {
  const { surface_water, diameter } = planet

  if (surface_water === 'unknown') return 'unknown'

  const radius = diameter / 2
  const area = 4 * Math.PI * radius * radius

  return area * surface_water * 0.01
}

function renderTable(data) {
  console.log('got data', data)

  const table = document.createElement('table')
  const tableHeader = document.createElement('thead')
  const headerRow = document.createElement('tr')
  const tableBody = document.createElement('tbody')

  rowHeaders.forEach((header) => {
    const th = document.createElement('th')
    th.appendChild(document.createTextNode(header))
    headerRow.appendChild(th)
  })

  tableHeader.appendChild(headerRow)
  table.appendChild(tableHeader)
  table.appendChild(tableBody)

  data.results
    .sort((a, b) => {
      const aName = a.name.toUpperCase()
      const bName = b.name.toUpperCase()

      if (aName < bName) return -1
      if (aName > bName) return 1
      return 0
    })
    .forEach((row) => {
      const tr = document.createElement('tr')

      rowDataFns.forEach((col) => {
        const td = document.createElement('td')
        td.appendChild(col(row))
        tr.appendChild(td)
      })

      tableBody.appendChild(tr)
    })

  document.body.innerHTML = ''
  document.body.appendChild(table)
}
