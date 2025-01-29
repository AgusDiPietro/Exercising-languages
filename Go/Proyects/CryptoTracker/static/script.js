document.addEventListener('DOMContentLoaded', function() {
    fetch('/api/top10')
      .then(response => response.json())
      .then(data => {
        const tableBody = document.querySelector('#cryptoTable tbody');
        data.forEach(coin => {
          const row = document.createElement('tr');
          row.innerHTML = `
            <td>${coin.name}</td>
            <td>${coin.symbol}</td>
            <td>${coin.current_price}</td>
            <td>${coin.market_cap}</td>
          `;
          tableBody.appendChild(row);
        });
      })
      .catch(error => console.error('Error fetching data:', error));
  });