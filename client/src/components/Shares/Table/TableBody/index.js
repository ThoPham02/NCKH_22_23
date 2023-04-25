const TableBody = ({ listKey, listItem }) => {
  return (
    <tbody>
      {listItem.map((item, index) => (
        <tr key={index}>
          <td>{index + 1}</td>
          {listKey.map((key, index) => (
            <td key={index}>{item[key]}</td>
          ))}
          <td>
            <button style={{borderTop:"none",borderLeft:"none",borderRight:"none",width:"70px"}}/*onClick={() => handleRegister(item)}*/>Đăng kí</button>
          </td>
        </tr>
      ))}
    </tbody>
  );

};

export default TableBody;
