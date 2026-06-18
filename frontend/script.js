const USER_URL = "http://localhost:8081";
const ORDER_URL = "http://localhost:8083";

async function registerUser() {
    try {
        const response = await fetch(`${USER_URL}/register`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                name: document.getElementById("name").value,
                email: document.getElementById("email").value,
                password: document.getElementById("password").value,
                role: "customer"
            })
        });

        const data = await response.json();

        if (!response.ok) {
            alert("Register Gagal: " + (data.message || "Terjadi kesalahan"));
            return;
        }

        alert("Register berhasil\nUser ID: " + data.user_id);
        window.location.href = "login.html";

    } catch (error) {
        console.error("Error saat register:", error);
        alert("Gagal terhubung ke server user-service. Pastikan backend sudah jalan.");
    }
}

async function login() {
    try {
        const response = await fetch(`${USER_URL}/login`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                email: document.getElementById("email").value,
                password: document.getElementById("password").value
            })
        });

        const data = await response.json();

        if (!response.ok) {
            alert("Login Gagal: " + (data.message || "Email atau password salah"));
            return;
        }

        // Simpan token jika login sukses
        if (data.token) {
            localStorage.setItem("token", data.token);
            alert("Login berhasil");
            window.location.href = "order.html";
        } else {
            alert("Login berhasil, tetapi token tidak ditemukan dari respon server.");
        }

    } catch (error) {
        console.error("Error saat login:", error);
        alert("Gagal terhubung ke server user-service.");
    }
}

async function createOrder() {
    const token = localStorage.getItem("token");
    if (!token) {
        alert("Anda belum login atau token kedaluwarsa. Silakan login kembali.");
        window.location.href = "login.html";
        return;
    }

    try {
        const response = await fetch(`${ORDER_URL}/order`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": "Bearer " + token
            },
            body: JSON.stringify({
                user_id: parseInt(document.getElementById("user_id").value) || 0,
                nama_barang: document.getElementById("nama_barang").value,
                berat: parseInt(document.getElementById("berat").value) || 0,
                dimensi: document.getElementById("dimensi").value,
                jenis: document.getElementById("jenis").value,
                alamat_pengirim: document.getElementById("alamat_pengirim").value,
                alamat_penerima: document.getElementById("alamat_penerima").value,
                nama_penerima: document.getElementById("nama_penerima").value,       // 🔥 TAMBAHKAN INI
                no_telp_penerima: document.getElementById("no_telp_penerima").value // 🔥 TAMBAHKAN INI
            })
        });

        const data = await response.json();

        if (!response.ok) {
            alert("Gagal membuat order: " + (data.message || "Periksa kembali data Anda"));
            return;
        }

        document.getElementById("result").innerHTML = `
            <h3>Order Berhasil</h3>
            <p>Resi: ${data.resi || '-'}</p>
            <p>Status: ${data.status || '-'}</p>
            <p>ETA: ${data.eta || '-'}</p>
        `;

    } catch (error) {
        console.error("Error saat membuat order:", error);
        alert("Gagal terhubung ke server order-service.");
    }
}