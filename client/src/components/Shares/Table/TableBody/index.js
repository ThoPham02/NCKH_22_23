const TableBody = ({ listKey, listItem }) => {
  const renderItem = (item) => {
    return listKey.map((key, index) => <td key={index}>{item[key]}</td>);
  };

  return (
    <tbody>
      {listItem.map((item, index) => {
        return <tr key={index}>{renderItem(item)}</tr>;
      })}
    </tbody>
  );
};

export default TableBody;
