CREATE TABLE product (
    product_id SERIAL PRIMARY KEY,         -- Auto-incremented primary key
    product_name VARCHAR(255) NOT NULL,     -- Product name, can't be NULL
    product_price INT NOT NULL,            -- Product price (integer, for currency in integer format)
    product_quality INT NOT NULL,          -- Product quality rating (integer)
    date_create TIMESTAMP DEFAULT CURRENT_TIMESTAMP,   -- Date and time of creation, defaults to the current timestamp
    date_update TIMESTAMP DEFAULT CURRENT_TIMESTAMP,   -- Date and time of last update, defaults to the current timestamp
    product_type VARCHAR(50),              -- Product type (e.g., Sedan, SUV, Pickup)
    brand VARCHAR(100),                    -- Brand name (e.g., Toyota, Honda, Mercedes-Benz)
    store VARCHAR(100),                    -- Store name (e.g., Timezone)
    status VARCHAR(50),                    -- Product status (e.g., regular, recommend)
    product_description TEXT,               -- Product description (long text)
    Image VARCHAR(1000),
    promotion_name TEXT,
    discount_money INT,
    discount_percent INT
);

CREATE TABLE cart (
    cart_id SERIAL PRIMARY KEY,                 -- รหัสตะกร้า
    product_id INT NOT NULL,                    -- รหัสสินค้า (เชื่อมโยงกับตารางสินค้า)
    quantity INT NOT NULL DEFAULT 1,            -- จำนวนสินค้า
    date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- วันที่เพิ่มสินค้า
    date_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- วันที่อัปเดตข้อมูล
    total_price DECIMAL(10, 2),                 -- ราคาสินค้าทั้งหมด
    store TEXT,                                 -- รหัสร้านค้าที่ขายสินค้า
    FOREIGN KEY (product_id) REFERENCES product(product_id)  -- เชื่อมโยงกับตารางสินค้า
);

CREATE TABLE comments (
    comment_id SERIAL PRIMARY KEY,  -- รหัสคอมเมนต์
    comment_text TEXT NOT NULL,     -- ข้อความคอมเมนต์
 rating INT ,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- เวลาที่คอมเมนต์ถูกสร้าง
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- เวลาที่คอมเมนต์ถูกแก้ไขล่าสุด
);



INSERT INTO product (product_id, product_name, product_price, product_quality, date_create, date_update, product_type, brand, store, status, product_description, Image, promotion_name, discount_money, discount_percent)
VALUES
(1, 'Camry', 1550000, 4, '2024-11-01 00:00:00', '2024-11-05 00:00:00', 'Sedan', 'Toyota', 'Timezone', 'regular', 'รถยนต์นั่งขนาดกลาง ถึงขนาดใหญ่ที่มาพร้อมเครื่องยนต์เบนซินและไฮบริด', 'https://www.toyota.co.th/media/product/series/grades/v/camry/185/5218fee9a08edf612a42103b76279f20ebcf87df6a2c49ec40957759191a25e3.webp', 'ลด 10%', 0, 10),
(2, 'Corolla', 750000, 2, '2024-11-01 00:00:00', '2024-11-05 00:00:00', 'Sedan', 'Toyota', 'Timeup', 'recommend', 'รถยนต์นั่งขนาดเล็กถึงขนาดกลาง มีตัวเลือกเครื่องยนต์เบนซินและไฮบริด', 'https://www.toyotakhonkaen.com/wp-content/uploads/2022/03/Corolla-Altis-GR-Sport_Platinum-White-Pearl.png', 'ลด 10%', 0, 10),
(3, 'Hilux', 900000, 5, '2024-11-02 00:00:00', '2024-11-05 00:00:00', 'Pickup', 'Toyota', 'Timemai', 'regular', 'รถกระบะขับเคลื่อน 4 ล้อที่แข็งแรง ทนทาน เหมาะกับงานหนักและการขับขี่ในสภาพถนนที่ยากลำบาก', 'https://www.toyota.co.th/media/product/series/colors/v/hilux_revo_prerunner/190/4268d5b4c83998ffd2bc2da26312c49135779829f52cf79ea7ca37ef3ddf9222.webp', 'ลด 10,000 บาท', 10000, 0),
(4, 'Civic', 1200000, 6, '2024-11-04 00:00:00', '2024-11-05 00:00:00', 'Sedan', 'Honda', 'Maitime', 'recommend', 'รถยนต์นั่งขนาดกลางที่มีดีไซน์ทันสมัยและสมรรถนะดีเยี่ยม', 'https://www.imc.co.th/img/model/640_2019110920183150.jpg', 'ลด 5%', 0, 5),
(5, 'Accord', 1500000, 7, '2024-10-30 00:00:00', '2024-11-05 00:00:00', 'Sedan', 'Honda', 'Timesleep', 'regular', 'รถยนต์นั่งขนาดกลางถึงขนาดใหญ่ที่มาพร้อมกับเครื่องยนต์ที่มีสมรรถนะสูง', 'https://www.imc.co.th/img/model/640_2023101721260536.jpg', 'ลด 10,000 บาท', 10000, 0),
(6, 'CR-V', 1500000, 6, '2024-11-01 00:00:00', '2024-11-05 00:00:00', 'SUV', 'Honda', 'Timezone', 'regular', 'SUV ขนาดกลางที่เหมาะสำหรับการใช้งานทั้งในเมืองและการขับขี่ในสถานการณ์ต่าง ๆ', 'https://images.dealer.com/ddc/vehicles/2024/Honda/CR-V/SUV/trim_EX_748dab/color/Canyon%20River%20Blue%20Metallic-BU-37%2C45%2C60-640-en_US.jpg?impolicy=downsize_bkpt&imdensity=1&w=520', 'ลด 10%', 0, 10),
(7, 'C-Class', 2400000, 3, '2024-11-03 00:00:00', '2024-11-05 00:00:00', 'Sedan', 'Mercedes-Benz', 'Timeup', 'recommend', 'รถยนต์นั่งหรูหราขนาดกลางพร้อมเครื่องยนต์เบนซิน', 'https://www.mercedes-benz.co.th/content/dam/hq/passengercars/cars/c-class/c-class-saloon-w206-pi/modeloverview/06-2022/images/mercedes-benz-c-class-w206-modeloverview-696x392-06-2022.png', 'ลด 10%', 0, 10),
(8, 'E-Class', 3500000, 3, '2024-11-03 00:00:00', '2024-11-05 00:00:00', 'Sedan', 'Mercedes-Benz', 'Timezone', 'regular', 'รถยนต์หรูขนาดกลางถึงใหญ่ มีทั้งเครื่องยนต์เบนซินและดีเซล', 'https://www.mercedes-benz.co.th/content/thailand/th/passengercars/models/saloon/w214-23-1/overview/_jcr_content/root/responsivegrid/media_slider_copy_co/media_slider_item/internal_video.component.damq1.3372250790823.jpg/mercedes-benz-e-class-w214-highlights-videostill-3302x1858-02-2023.jpg', 'ลด 15%', 0, 15),
(9, 'GLC', 3000000, 4, '2024-10-31 00:00:00', '2024-11-05 00:00:00', 'SUV', 'Mercedes-Benz', 'Maitime', 'regular', 'SUV ขนาดกลางที่มาพร้อมกับความหรูหราและเทคโนโลยีสูง', 'https://www.mercedes-benz.co.th/content/dam/hq/passengercars/cars/glc/suv-x254/modeloverview/images/mercedes-benz-glc-suv-x254-modeloverview-696x392-05-2022.png', 'ลด 15,000 บาท', 15000, 0),
(10, '3 Series', 2200000, 4, '2024-10-31 00:00:00', '2024-11-05 00:00:00', 'Sedan', 'BMW', 'Timesleep', 'recommend', 'รถยนต์นั่งหรูขนาดกลางที่มาพร้อมกับสมรรถนะที่ยอดเยี่ยม', 'https://www.bmw.co.th/content/dam/bmw/common/all-models/3-series/sedan/2022/navigation/bmw-3-series-sedan-lci-modelfinder.png', 'ลด 20,000 บาท', 20000, 0),
(11, '5 Series', 3200000, 3, '2024-11-01 00:00:00', '2024-11-05 00:00:00', 'Sedan', 'BMW', 'Timezone', 'regular', 'รถยนต์หรูขนาดกลางถึงใหญ่ มาพร้อมกับเทคโนโลยีทันสมัยและสมรรถนะสูง', 'https://www.bmw.co.th/content/dam/bmw/common/all-models/5-series/sedan/2023/5-series-sedan-silver.png', 'ลด 10%', 0, 10),
(12, 'X5', 4200000, 4, '2024-10-28 00:00:00', '2024-11-05 00:00:00', 'SUV', 'BMW', 'Timeup', 'regular', 'SUV หรูขนาดใหญ่ที่มาพร้อมกับความปลอดภัยและสมรรถนะที่ยอดเยี่ยม', 'https://www.imc.co.th/img/model/640_2019120111181447.jpg', 'ลด 15%', 0, 15),
(13, 'Elantra', 900000, 6, '2024-10-23 00:00:00', '2024-11-05 00:00:00', 'Sedan', 'Hyundai', 'Timemai', 'regular', 'รถยนต์นั่งขนาดกลางที่มาพร้อมเทคโนโลยีที่ทันสมัยและประหยัดน้ำมัน', 'https://5.imimg.com/data5/YH/MM/MY-13414525/hyundai-elantra-car-500x500.png', 'ลด 10,000 บาท', 10000, 0),
(14, 'Tucson', 1200000, 3, '2024-11-01 00:00:00', '2024-11-05 00:00:00', 'SUV', 'Hyundai', 'Maitime', 'regular', 'SUV ขนาดกลางที่มาพร้อมกับเครื่องยนต์เบนซินและไฮบริด', 'https://cdn.wheel-size.com/thumbs/30/ff/30ffb5b03f2254dffe090918d5abea1c.jpg', 'ลด 15,000 บาท', 15000, 0),
(15, 'Sonata', 1200000, 3, '2024-10-23 00:00:00', '2024-11-05 00:00:00', 'Sedan', 'Hyundai', 'Timesleep', 'recommend', 'รถยนต์นั่งขนาดกลางที่มาพร้อมเทคโนโลยีที่ทันสมัยและระบบความปลอดภัยขั้นสูง', 'https://cdn.wheel-size.com/automobile/body/hyundai-sonata-2014-2017-1640670465.7280176.jpg', 'ลด 15,000 บาท', 15000, 0),
(16, 'A4', 2200000, 3, '2024-10-29 00:00:00', '2024-11-05 00:00:00', 'Sedan', 'Audi', 'Timezone', 'recommend', 'รถยนต์นั่งหรูขนาดกลางที่มาพร้อมสมรรถนะสูงและเทคโนโลยีล้ำสมัย', 'https://d2ivfcfbdvj3sm.cloudfront.net/7fc965ab77efe6e0fa62e4ca1ea7673bb25f44570f1e3d8e88cb10/stills_0640_png/MY2022/15564/15564_st0640_116.png', 'ลด 10%', 0, 10),
(17, 'Q5', 3000000, 4, '2024-11-01 00:00:00', '2024-11-05 00:00:00', 'SUV', 'Audi', 'Timeup', 'regular', 'SUV ขนาดกลางที่หรูหราและมาพร้อมเทคโนโลยีการขับขี่ขั้นสูง', 'https://mediaservice.audi.com/media/cdb/data/1a89eac3-e29b-4b00-8dc9-7a2e11cc64d7.jpg', 'ลด 10%', 0, 10),
(18, 'A6', 3200000, 4, '2024-11-03 00:00:00', '2024-11-05 00:00:00', 'Sedan', 'Audi', 'Timemai', 'regular', 'รถยนต์นั่งหรูขนาดใหญ่ที่มีระบบขับขี่ที่มีประสิทธิภาพและความปลอดภัยสูง', 'https://cdn.wheel-size.com/thumbs/2c/de/2cdeec40d1af3b181592c50d5b993ff5.jpg', 'ลด 15,000 บาท', 15000, 0),
(19, 'Yaris', 550000, 3, '2024-11-01 00:00:00', '2024-11-05 00:00:00', 'Hatchback', 'Toyota', 'Maitime', 'regular', 'รถยนต์ขนาดเล็กที่ประหยัดน้ำมัน เหมาะสำหรับการขับขี่ในเมือง', 'https://www.checkraka.com/uploaded/_resize/max350x197/34/342c2df1415752c7b0a296224f75b83d.webp', 'ลด 15,000 บาท', 15000, 0),
(20, 'Jazz', 650000, 4, '2024-11-02 00:00:00', '2024-11-05 00:00:00', 'Hatchback', 'Honda', 'Timesleep', 'recommend', 'รถยนต์แฮทช์แบคขนาดเล็กที่มีการออกแบบที่สวยงามและพื้นที่ภายในที่กว้างขวาง', 'https://www.hongtonggas.com/wp-content/uploads/2022/08/Honda-Jazz.jpg', 'ลด 20%', 0, 20),
(21, 'Fortuner', 1800000, 5, '2024-11-01 00:00:00', '2024-11-05 00:00:00', 'SUV', 'Toyota', 'Timesleep', 'regular', 'SUV ขนาดใหญ่ที่แข็งแรงและเหมาะสำหรับทุกสภาพถนน', 'https://toyotabuzz.com/img/upload/car/small/20220801151935_782595128.png', 'ลด 15,000 บาท', 15000, 0),
(22, 'HR-V', 1100000, 4, '2024-11-03 00:00:00', '2024-11-05 00:00:00', 'SUV', 'Honda', 'Timeup', 'recommend', 'SUV ขนาดกลางที่มีการออกแบบสไตล์ทันสมัยและประหยัดน้ำมัน', 'https://api.hwcapis.com:12443/PRD.VDS.API.CAR/Media/view/1e56147096804f289a207544dae05d72', 'ลด 20,000 บาท', 20000, 0),
(23, 'GLE', 4800000, 5, '2024-10-31 00:00:00', '2024-11-05 00:00:00', 'SUV', 'Mercedes-Benz', 'Timemai', 'regular', 'SUV ขนาดใหญ่ที่หรูหราและมาพร้อมเทคโนโลยีขั้นสูง', 'https://www.mercedes-benz.co.th/content/dam/hq/passengercars/cars/gle/gle-suv-v167-fl-pi/modeloverview/01-2023/images/mercedes-benz-gle-suv-v167-modeloverview-696x392-01-2023.png', 'ลด 10,000 บาท', 10000, 0),
(24, 'X3', 3200000, 4, '2024-10-28 00:00:00', '2024-11-05 00:00:00', 'SUV', 'BMW', 'Maitime', 'recommend', 'SUV ขนาดกลางที่มีสมรรถนะยอดเยี่ยมและความสะดวกสบาย', 'https://www.bmw.co.th/content/dam/bmw/marketTH/common/All%20Models/x-series/x3/2021/BMW%20X3%20xDrive30e.png', 'ลด 15,000 บาท', 15000, 0),
(25, 'Q3', 2500000, 3, '2024-11-02 00:00:00', '2024-11-05 00:00:00', 'SUV', 'Audi', 'Timesleep', 'regular', 'SUV ขนาดเล็กที่หรูหราและสะดวกสบายสำหรับการขับขี่ในเมือง', 'https://mediaservice.audi.com/media/cdb/data/22148bd8-ec0a-4ce4-b702-25fada0fb8ef.jpg', 'ลด 15%', 0, 15),
(26, 'Altima', 950000, 4, '2024-11-03 00:00:00', '2024-11-05 00:00:00', 'Sedan', 'Nissan', 'Timezone', 'recommend', 'รถยนต์นั่งขนาดกลางที่มีความประหยัดน้ำมันและเทคโนโลยีทันสมัย', 'https://vehicle-images.dealerinspire.com/stock-images/chrome/cf2137d64cf96bd2f6e00749abe83a36.png', 'ลด 10,000 บาท', 10000, 0),
(27, 'CX-5', 1600000, 5, '2024-11-04 00:00:00', '2024-11-05 00:00:00', 'SUV', 'Mazda', 'Timeup', 'regular', 'SUV ขนาดกลางที่มีความคล่องตัวและการออกแบบที่หรูหรา', 'https://www.imc.co.th/img/model/640_2021101121475511.jpg', 'ลด 15,000 บาท', 15000, 0),
(28, 'Sentra', 800000, 3, '2024-11-01 00:00:00', '2024-11-05 00:00:00', 'Sedan', 'Nissan', 'Timemai', 'regular', 'รถยนต์นั่งขนาดกลางที่มีความปลอดภัยสูงและประหยัดน้ำมัน', 'https://di-uploads-pod35.dealerinspire.com/newtonnissanofgallatin/uploads/2021/02/2021-Nissan-Sentra-S-Model-Left.jpg', 'ลด 15%', 0, 15),
(29, 'Optima', 1300000, 3, '2024-11-02 00:00:00', '2024-11-05 00:00:00', 'Sedan', 'Kia', 'Maitime', 'regular', 'รถยนต์นั่งขนาดกลางที่มาพร้อมกับสมรรถนะและความปลอดภัยที่ดีเยี่ยม', 'https://crdms.images.consumerreports.org/c_lfill,w_470,q_auto,f_auto/prod/cars/chrome/white/2020KIC050033_1280_01', 'ลด 10%', 0, 10),

(34, 'Model Y', 2500000, 5, '2024-11-04 00:00:00', '2024-11-05 00:00:00', 'SUV', 'Tesla', 'Maitime', 'recommend', 'SUV ไฟฟ้าที่มาพร้อมกับการออกแบบที่ทันสมัยและเทคโนโลยีชั้นนำ', 'https://www.tesla.com/ownersmanual/images/GUID-1F2D8746-336F-4CF9-9A04-F35E960F31FE-online-en-US.png', 'ลด 15,000 บาท', 15000, 0),
(35, 'Leaf', 900000, 3, '2024-11-01 00:00:00', '2024-11-05 00:00:00', 'Hatchback', 'Nissan', 'Timezone', 'regular', 'รถยนต์ไฟฟ้าขนาดเล็กที่เหมาะสำหรับการใช้งานในเมือง', 'https://img.pcauto.com/model/images/touPic/th/Nissan-Leaf_1318.png', 'ลด 10%', 0, 10),
(36, 'Mustang', 3200000, 4, '2024-10-31 00:00:00', '2024-11-05 00:00:00', 'Coupe', 'Ford', 'Timeup', 'regular', 'รถยนต์สปอร์ตคูเป้ที่มาพร้อมกับสมรรถนะที่ยอดเยี่ยมและดีไซน์ที่เป็นเอกลักษณ์', 'https://autoinfo.co.th/uploads/2021/08/Ford-Mustang-carbonized-grey-w.jpg', 'ลด 10%', 0, 10),
(37, 'Explorer', 3600000, 4, '2024-10-29 00:00:00', '2024-11-05 00:00:00', 'SUV', 'Ford', 'Timemai', 'recommend', 'SUV ขนาดใหญ่ที่เหมาะสำหรับครอบครัวและมาพร้อมกับเทคโนโลยีทันสมัย', 'https://d2qldpouxvc097.cloudfront.net/image-by-path?bucket=a5-gallery-serverless-prod-chromebucket-1iz9ffi08lwxm&key=450239%2Ffront34%2Flg%2Fe7e7e7', 'ลด 15%', 0, 15),
(38, 'Outlander', 1500000, 3, '2024-11-03 00:00:00', '2024-11-05 00:00:00', 'SUV', 'Mitsubishi', 'Maitime', 'regular', 'SUV ขนาดกลางที่มีความทนทานและเหมาะสำหรับการใช้งานทุกสภาพถนน', 'https://f.ptcdn.info/594/072/000/qonya9zoh1OJ0X1hTJM-o.jpg', 'ลด 15%', 0, 15),
(39, 'Tucson', 1100000, 2, '2024-11-03 00:00:00', '2024-11-05 00:00:00', 'SUV', 'Hyundai', 'Timezone', 'regular', 'SUV ขนาดกลางที่มาพร้อมกับเครื่องยนต์เบนซินและไฮบริด', 'https://cdn.wheel-size.com/thumbs/30/ff/30ffb5b03f2254dffe090918d5abea1c.jpg', 'ลด 10%', 0, 10),
(40, 'Sonata', 1100000, 3, '2024-11-03 00:00:00', '2024-11-05 00:00:00', 'Sedan', 'Hyundai', 'Timeup', 'regular', 'รถยนต์นั่งขนาดกลางที่มาพร้อมเทคโนโลยีที่ทันสมัยและระบบความปลอดภัยขั้นสูง', 'https://cdn.wheel-size.com/automobile/body/hyundai-sonata-2014-2017-1640670465.7280176.jpg', 'ลด 10%', 0, 10),
(41, '3 Series', 2100000, 2, '2024-10-31 00:00:00', '2024-11-05 00:00:00', 'Sedan', 'BMW', 'Timemai', 'recommend', 'รถยนต์นั่งหรูขนาดกลางที่มาพร้อมกับสมรรถนะที่ยอดเยี่ยม', 'https://www.bmw.co.th/content/dam/bmw/common/all-models/3-series/sedan/2022/navigation/bmw-3-series-sedan-lci-modelfinder.png', 'ลด 10%', 0, 10),
(42, 'Civic', 1300000, 3, '2024-11-02 00:00:00', '2024-11-05 00:00:00', 'Sedan', 'Honda', 'Timemai', 'recommend', 'รถยนต์นั่งขนาดกลางที่มีดีไซน์ทันสมัยและสมรรถนะดีเยี่ยม', 'https://www.imc.co.th/img/model/640_2019110920183150.jpg', 'ลด 10%', 0, 10),
(43, 'GLC', 3000000, 3, '2024-11-03 00:00:00', '2024-11-05 00:00:00', 'SUV', 'Mercedes-Benz', 'Timesleep', 'regular', 'SUV ขนาดกลางที่มาพร้อมกับความหรูหราและเทคโนโลยีสูง', 'https://www.mercedes-benz.co.th/content/dam/hq/passengercars/cars/glc/suv-x254/modeloverview/images/mercedes-benz-glc-suv-x254-modeloverview-696x392-05-2022.png', 'ลด 10%', 0, 10),
(44, 'Q5', 3000000, 2, '2024-11-03 00:00:00', '2024-11-05 00:00:00', 'SUV', 'Audi', 'Maitime', 'recommend', 'SUV ขนาดกลางที่หรูหราและมาพร้อมเทคโนโลยีการขับขี่ขั้นสูง', 'https://mediaservice.audi.com/media/cdb/data/1a89eac3-e29b-4b00-8dc9-7a2e11cc64d7.jpg', 'ลด 15%', 0, 15);
