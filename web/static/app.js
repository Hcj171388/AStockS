const cards = document.querySelector('#cards');
const refresh = document.querySelector('#refresh');

function money(value) {
  const abs = Math.abs(value);
  if (abs >= 100000000) return `${(value / 100000000).toFixed(2)}亿`;
  if (abs >= 10000) return `${(value / 10000).toFixed(2)}万`;
  return value.toFixed(2);
}

function pct(value) { return `${(value * 100).toFixed(2)}%`; }

function render(stocks) {
  cards.innerHTML = stocks.map((stock, index) => `
    <article class="card">
      <div class="rank"><div><strong>TOP ${index + 1}</strong><h3>${stock.name}</h3><p class="code">${stock.code}</p></div><div class="score">${stock.score}</div></div>
      <div class="metrics">
        <div class="metric"><span>资金净流入</span><b>${money(stock.netInflow)}</b></div>
        <div class="metric"><span>资金强度 F</span><b>${pct(stock.fundStrength)}</b></div>
        <div class="metric"><span>资金层得分</span><b>${stock.fundScore}</b></div>
        <div class="metric"><span>价格层得分</span><b>${stock.priceScore}</b></div>
        <div class="metric"><span>筹码层得分</span><b>${stock.chipScore}</b></div>
        <div class="metric"><span>连续流入</span><b>${stock.consecutiveInflowDays} 天</b></div>
      </div>
      <div class="tags">${stock.reasons.map(reason => `<span>${reason}</span>`).join('')}</div>
    </article>`).join('');
}

async function loadRank() {
  refresh.disabled = true;
  refresh.textContent = '加载中...';
  try {
    const response = await fetch('/api/rank');
    if (!response.ok) throw new Error(`HTTP ${response.status}`);
    render(await response.json());
  } finally {
    refresh.disabled = false;
    refresh.textContent = '刷新评分';
  }
}

refresh.addEventListener('click', loadRank);
loadRank();
