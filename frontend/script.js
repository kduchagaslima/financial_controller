async function loadSummary() {
  const res = await fetch('/categories/summary');
  const data = await res.json();
  const container = document.getElementById('summary');
  container.innerHTML = '';
  data.forEach(item => {
    const div = document.createElement('div');
    div.className = 'category';
    div.innerHTML = `<strong>${item.category}</strong>: $${item.total.toFixed(2)} <button onclick="toggle(${item.category})">Toggle</button>`;
    const transactions = document.createElement('div');
    transactions.id = 'cat-' + item.category;
    transactions.className = 'transactions';
    transactions.innerHTML = '<button onclick="addTransaction(\'' + item.category + '\')">Add</button>';
    div.appendChild(transactions);
    container.appendChild(div);
  });
}

function toggle(cat) {
  const t = document.getElementById('cat-' + cat);
  if (t.style.display === 'block') t.style.display = 'none';
  else t.style.display = 'block';
}

async function addTransaction(categoryName) {
  const description = prompt('Description');
  const amount = parseFloat(prompt('Amount'));
  const dateStr = prompt('Date (YYYY-MM-DD)');
  const categoryIdRes = await fetch(`/categories/summary`);
  const categories = await categoryIdRes.json();
  const catObj = categories.find(c => c.category === categoryName);
  const resp = await fetch('/transactions', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ description, amount, date: dateStr, category_id: catObj.id })
  });
  if (resp.ok) loadSummary();
}

async function loadReport() {
  const catId = document.getElementById('reportCategory').value;
  const start = document.getElementById('reportStart').value;
  const end = document.getElementById('reportEnd').value;
  const params = new URLSearchParams();
  if (catId) params.append('category_id', catId);
  if (start) params.append('start', start);
  if (end) params.append('end', end);
  const res = await fetch('/reports?' + params.toString());
  const data = await res.json();
  document.getElementById('reportResult').textContent = JSON.stringify(data, null, 2);
}

loadSummary();
