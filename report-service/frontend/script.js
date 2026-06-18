const BASE_URL = "http://localhost:8083";

async function loadDailyReport() {

    const response = await fetch(
        `${BASE_URL}/report/daily?date=2026-06-18`
    );

    const data = await response.json();

    document.getElementById("dailyReport").innerHTML = `
        <div class="card">
            <h3>Daily Report</h3>
            <p>Total Paket: ${data.total_paket}</p>
            <p>Delivered: ${data.delivered}</p>
            <p>Pending: ${data.pending}</p>
            <p>Terlambat: ${data.terlambat}</p>
            <p>Rata-rata ETA: ${data.rata_rata_eta}</p>
        </div>
    `;
}

async function loadProblems() {

    const response = await fetch(
        `${BASE_URL}/report/problems`
    );

    const data = await response.json();

    let html = `
        <div class="card">
        <h3>Problem Packages</h3>
    `;

    data.forEach(item => {
        html += `
            <p>${item.resi} - ${item.status}</p>
        `;
    });

    html += "</div>";

    document.getElementById("problems").innerHTML = html;
}

async function loadCourierPerformance() {

    const response = await fetch(
        `${BASE_URL}/report/courier-performance`
    );

    const data = await response.json();

    document.getElementById("courier").innerHTML = `
        <div class="card">
            <h3>Courier Performance</h3>
            <p>Courier ID: ${data.courier_id}</p>
            <p>Total Pengiriman: ${data.total_pengiriman}</p>
            <p>Berhasil: ${data.berhasil}</p>
            <p>Terlambat: ${data.terlambat}</p>
            <p>Score: ${data.score}</p>
        </div>
    `;
}