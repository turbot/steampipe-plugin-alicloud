select name, org_name, issuer, buy_in_aliyun, common, cert, key
from alicloud_cas_certificate
where id = '{{ output.certificate_id.value }}';