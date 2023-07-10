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
$sql = "SELECT * FROM users WHERE username = '$username' AND password = '$password'"; // แก้ไขตามตารางและฟิลด์ที่เกี่ยวข้อง

// ส่งคำสั่ง SQL ไปยังฐานข้อมูล
$result = mysqli_query($conn, $sql);

// ตรวจสอบผลลัพธ์
if (mysqli_num_rows($result) > 0) {
    echo "ล็อคอินสำเร็จ";
} else {
    echo "ไม่พบชื่อผู้ใช้งานและรหัสผ่าน";
}

// ปิดการเชื่อมต่อฐานข้อมูล
mysqli_close($conn);
