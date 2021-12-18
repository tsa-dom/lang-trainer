import React from 'react'
import { Container, Table } from 'react-bootstrap'

// eslint-disable-next-line no-unused-vars
const ItemList = ({ rows, columns, onCellClick, handleItemRemove }) => {

  return (
    <Container>
      <Table striped bordered hover>
        <thead>
          <tr>
            {columns.map(col => <th key={col.field}>{col.headerName}</th>)}
          </tr>
        </thead>
        <tbody>
          {rows.map(row => <tr onClick={() => onCellClick(row)} style={{ cursor: 'pointer' }} key={row.id}>
            {columns.map(col => <td key={col.field}>{row[col.field]}</td>)}
          </tr>)}
        </tbody>
      </Table>
    </Container>
  )
}

export default ItemList