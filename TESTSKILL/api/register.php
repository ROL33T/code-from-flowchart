<?php
$servername = "127.0.0.1";
$username = "root";
$password = "";
$dbname = "test_skill";

// เชื่อมต่อฐานข้อมูล MySQL
$conn = mysqli_connect($servername, $username, $password, $dbname);

// ตรวจสอบการเชื่อมต่อฐานข้อมูล
if (!$conn) {
    die("การเชื่อมต่อฐานข้อมูลล้มเหลว: " . mysqli_connect_error());
}

// รับค่าจากแบบฟอร์ม
$username = $_POST["username"];
$password = $_POST["password"];

// สร้างคำสั่ง SQL สำหรับการเลือกข้อมูล
$sql = "SELECT * FROM users WHERE username = '$username' OR password = '$password'"; // แก้ไขตามตารางและฟิลด์ที่เกี่ยวข้อง

// ส่งคำสั่ง SQL ไปยังฐานข้อมูล
$result = mysqli_query($conn, $sql);

// ตรวจสอบผลลัพธ์
if (mysqli_num_rows($result) > 0) {
    // มีผลลัพธ์
    echo "ไม่สามารถสมัครได้เนื่องจากซ้ำ";
} else {
    // ไม่พบผลลัพธ์
    $sql = "INSERT INTO users (username, password) VALUES ('$username', '$password')"; // แก้ไขตามตารางและฟิลด์ที่เกี่ยวข้อง
    $result = mysqli_query($conn, $sql);
    if ($result) {
        echo "REGISTER OK";
    } else {
        echo "เกิดข้อผิดพลาดในการสมัคร";
    }
}

// ปิดการเชื่อมต่อฐานข้อมูล
mysqli_close($conn);
