-- Copyright 2019 Benjamin Böhmke <benjamin@boehmke.net>.
--
-- Licensed under the Apache License, Version 2.0 (the "License");
-- you may not use this file except in compliance with the License.
-- You may obtain a copy of the License at
--
--     http://www.apache.org/licenses/LICENSE-2.0
--
-- Unless required by applicable law or agreed to in writing, software
-- distributed under the License is distributed on an "AS IS" BASIS,
-- WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
-- See the License for the specific language governing permissions and
-- limitations under the License.

local spw_proto = Proto("spw","SMA Speedwire")


local pf_spw_id = ProtoField.new("ID", "spw.id", ftypes.STRING)
local pf_spw_end = ProtoField.uint32("spw.end", "End", base.HEX)

-- general fields
local pf_spw_entry_length = ProtoField.uint16("spw.entry_length", "Entry Length", base.HEX)
local ENTRY_TAGS = {
    [0x02a0] = "Tag0 (Group)",
    [0x0010] = "SMA Net 2",
    [0x0020] = "Discovery",
    [0x0030] = "IP Address",
}
local pf_spw_entry_tag = ProtoField.uint16("spw.entry_tag", "Entry", base.HEX, ENTRY_TAGS)
local pf_spw_raw = ProtoField.new("Raw data", "spw.raw", ftypes.BYTES)

-- Tag0 (Group)
local pf_spw_tag0_group = ProtoField.uint32("spw.tag0.group", "Group", base.HEX)

-- IP Address
local pf_spw_ip = ProtoField.ipv4("spw.ip", "IP", base.HEX)

-- # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
-- SMA Net 2
-- # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
local SMA_NET2_IDS = {
    [0x6069] = "Energy meter",
    [0x6065] = "Device data",
}
local pf_spw_net2_id = ProtoField.uint16("spw.net2.id", "Protocol ID", base.HEX, SMA_NET2_IDS)

-- Energy meter
local pf_net2_em_susys = ProtoField.uint16("spw.net2.em.susys_id", "SusyID", base.DEC)
local pf_net2_em_ser_no = ProtoField.uint32("spw.net2.em.ser_no", "SerNo", base.DEC)
local pf_net2_em_ticker = ProtoField.uint32("spw.net2.em.ticker", "Ticker (ms)", base.DEC)

local pf_net2_em_obis_chan = ProtoField.uint8("spw.net2.em.obis.chan", "Chanel", base.DEC)
local pf_net2_em_obis_value = ProtoField.uint8("spw.net2.em.obis.value", "value", base.DEC)
local pf_net2_em_obis_type = ProtoField.uint8("spw.net2.em.obis.type", "type", base.DEC)
local pf_net2_em_obis_tariff = ProtoField.uint8("spw.net2.em.obis.tariff", "Tariff", base.DEC)

local pf_net2_em_val = ProtoField.uint32("spw.net2.em.value", "Value", base.DEC)
local pf_net2_em_val_l = ProtoField.uint64("spw.net2.em.value_large", "Value", base.DEC)

-- Device data
local pf_net2_dd_len = ProtoField.uint8("spw.net2.dd.length", "Length (1/4)", base.DEC)
local pf_net2_dd_control = ProtoField.uint8("spw.net2.dd.control", "Control", base.HEX)

local pf_net2_dd_dst_susys = ProtoField.uint16("spw.net2.dd.dst_susys_id", "Dst SusyID", base.DEC)
local pf_net2_dd_dst_ser_no = ProtoField.uint32("spw.net2.dd.dst_ser_no", "Dst SerNo", base.DEC)
local pf_net2_dd_job_num = ProtoField.uint8("spw.net2.dd.job_num", "Job number", base.HEX)
local pf_net2_dd_src_susys = ProtoField.uint16("spw.net2.dd.src_susys_id", "Src SusyID", base.DEC)
local pf_net2_dd_src_ser_no = ProtoField.uint32("spw.net2.dd.src_ser_no", "Src SerNo", base.DEC)

local pf_net2_dd_status = ProtoField.uint16("spw.net2.dd.status", "Status", base.HEX)
local pf_net2_dd_packet_count = ProtoField.uint16("spw.net2.dd.packet_count", "Packet count", base.HEX)
local pf_net2_dd_packet_id = ProtoField.uint16("spw.net2.dd.packet_id", "Packet ID", base.HEX)

local pf_net2_dd_command = ProtoField.uint8("spw.net2.dd.command", "Command", base.HEX)
local pf_net2_dd_object = ProtoField.uint16("spw.net2.dd.object", "Object", base.HEX)
local pf_net2_dd_param = ProtoField.uint32("spw.net2.dd.param", "Param", base.HEX)

local pf_net2_dd_resp_class = ProtoField.uint8("spw.net2.dd.resp_class", "Class", base.HEX)
local pf_net2_dd_resp_code = ProtoField.uint16("spw.net2.dd.resp_code", "Code", base.HEX)
local pf_net2_dd_resp_type = ProtoField.uint8("spw.net2.dd.resp_type", "Type", base.HEX)
local pf_net2_dd_resp_timestamp = ProtoField.uint64("spw.net2.dd.resp_timestamp", "Timestamp", base.DEC)

local pf_net2_dd_resp_value = ProtoField.uint32("spw.net2.dd.resp_value", "Val32", base.DEC)
local pf_net2_dd_resp_value_i = ProtoField.int32("spw.net2.dd.resp_value", "iVal32", base.DEC)
local pf_net2_dd_resp_value_l = ProtoField.uint64("spw.net2.dd.resp_value_l", "Val64", base.DEC)
local pf_net2_dd_resp_value_s = ProtoField.string("spw.net2.dd.resp_value_s", "ValStr")


spw_proto.fields = {
    pf_spw_id, pf_spw_end,
    pf_spw_entry_length, pf_spw_entry_tag, pf_spw_raw, 

    pf_spw_tag0_group,
    pf_spw_ip,

    pf_spw_net2_id,
    pf_net2_em_susys, pf_net2_em_ser_no, pf_net2_em_ticker,
    pf_net2_em_obis_chan, pf_net2_em_obis_value, pf_net2_em_obis_type, pf_net2_em_obis_tariff,
    pf_net2_em_val, pf_net2_em_val_l,

    pf_net2_dd_len, pf_net2_dd_control, 
    pf_net2_dd_dst_susys, pf_net2_dd_dst_ser_no, pf_net2_dd_job_num, pf_net2_dd_src_susys, pf_net2_dd_src_ser_no,
    pf_net2_dd_status, pf_net2_dd_packet_count, pf_net2_dd_packet_id,
    pf_net2_dd_command, pf_net2_dd_object, pf_net2_dd_param,

    pf_net2_dd_resp_class, pf_net2_dd_resp_code, pf_net2_dd_resp_type, pf_net2_dd_resp_timestamp,
    pf_net2_dd_resp_value, pf_net2_dd_resp_value_l, pf_net2_dd_resp_value_s, pf_net2_dd_resp_value_i,
}

-- dissect protocol
function spw_proto.dissector(buffer, pinfo, root)
    pinfo.cols.protocol = "Speedwire"

    local spw_tree = root:add(spw_proto, buffer, "")
    spw_tree:add(pf_spw_id, buffer:range(0x0, 4))

    local offset = 0x04
    local length = buffer:range(offset, 2):uint()
    local tag = buffer:range(offset+2, 2):uint()

    while tag > 0 or length > 0 do
        local entry_tree = spw_tree:add(spw_proto, buffer:range(offset, length+4), "Entry: "..entry_name(tag))
        entry_tree:add(pf_spw_entry_length, buffer:range(offset, 2))
        --entry_tree:add(pf_spw_entry_tag, buffer:range(offset+2, 2))
        
        local content = buffer:range(offset+4, length)
        -- Tag0 (Group)
        if tag == 0x02a0 then
            entry_tree:add(pf_spw_tag0_group, content)

        -- SMA Net 2
        elseif tag == 0x0010 then
            parse_net2(entry_tree, content)

        -- Discovery
        elseif tag == 0x0020 then
            -- No content

        -- IP Address
        elseif tag == 0x0030 then
            entry_tree:add(pf_spw_ip, content)

        else
            entry_tree:add(pf_spw_raw, content)
        end

        offset = offset + length + 4
        length = buffer:range(offset,2):uint()
        tag = buffer:range(offset+2, 2):uint()
    end
    spw_tree:add(pf_spw_end, buffer:range(offset, 4))
end

function entry_name(tag)
    local tag_string = string.format(" (0x%04X)", tag)
    if ENTRY_TAGS[tag] == nil then
        return "Unknown"..tag_string
    else
        return ENTRY_TAGS[tag]..tag_string
    end
end

-- load the udp.port table
udp_table = DissectorTable.get("udp.port")

-- register our protocol to handle udp port 12010
udp_table:add(9522, spw_proto)


-- SMA Net 2
function parse_net2(root, buffer)
    local id = buffer:range(0, 2):uint()
    local proto_tree = root:add(spw_proto, buffer, "Protocol: "..protocol_name(id))
    
    -- Energy meter
    local content = buffer:range(2)
    if id == 0x6069 then
        parse_net2_emergy_meter(proto_tree, content)

    -- device data
    elseif id == 0x6065 then
        parse_net2_device_data(proto_tree, content)

    else
        proto_tree:add(pf_spw_raw, content)
    end
end

function protocol_name(id)
    local id_string = string.format(" (0x%04X)", id)
    if SMA_NET2_IDS[id] == nil then
        return "Unknown"..id_string
    else
        return SMA_NET2_IDS[id]..id_string
    end
end


function parse_net2_emergy_meter(root, buffer)
    root:add(pf_net2_em_susys, buffer:range(0, 2))
    root:add(pf_net2_em_ser_no, buffer:range(2, 4))
    root:add(pf_net2_em_ticker, buffer:range(6, 4))

    local offset = 10
    while buffer:len() > offset+8 do
        local obis_str = string.format("%d:%d.%d.%d", 
            buffer:range(offset, 1):uint(), 
            buffer:range(offset+1, 1):uint(), 
            buffer:range(offset+2, 1):uint(), 
            buffer:range(offset+3, 1):uint())
        local value_str = ""

        -- obis type
        local val_length = 0
        if buffer:range(offset+2, 1):uint() == 8 then
            -- TODO find better solution for big integers
            value_str = string.format("%d", buffer:range(offset+9, 4):uint())
            val_length = 8
        else
            value_str = string.format("%d", buffer:range(offset+5, 4):uint())
            val_length = 4
        end

        local value_tree = root:add(spw_proto, buffer:range(offset, val_length+4), obis_str.." - "..value_str)
        local obis_tree = value_tree:add(spw_proto, buffer:range(offset, 4), obis_str)
        obis_tree:add(pf_net2_em_obis_chan, buffer:range(offset, 1))
        obis_tree:add(pf_net2_em_obis_value, buffer:range(offset+1, 1))
        obis_tree:add(pf_net2_em_obis_type, buffer:range(offset+2, 1))
        obis_tree:add(pf_net2_em_obis_tariff, buffer:range(offset+3, 1))

        if buffer:range(offset+2, 1):uint() == 8 then
            value_tree:add(pf_net2_em_val_l, buffer:range(offset+5, 8))
        else
            value_tree:add(pf_net2_em_val, buffer:range(offset+5, 4))
        end
        offset = offset + 4 + val_length
    end
end

function parse_net2_device_data(root, buffer)
    root:add_le(pf_net2_dd_len, buffer:range(0, 1))
    root:add_le(pf_net2_dd_control, buffer:range(1, 1))

    root:add_le(pf_net2_dd_dst_susys, buffer:range(2, 2))
    root:add_le(pf_net2_dd_dst_ser_no, buffer:range(4, 4))
    -- 8 = unknown
    root:add_le(pf_net2_dd_job_num, buffer:range(9, 1))
    root:add_le(pf_net2_dd_src_susys, buffer:range(10, 2))
    root:add_le(pf_net2_dd_src_ser_no, buffer:range(12, 4))
    -- 16,17 == 8,9 ???

    root:add_le(pf_net2_dd_status, buffer:range(18, 2))
    root:add_le(pf_net2_dd_packet_count, buffer:range(20, 2))
    root:add_le(pf_net2_dd_packet_id, buffer:range(22, 2))
    root:add_le(pf_net2_dd_command, buffer:range(24, 1))
    root:add_le(pf_net2_dd_object, buffer:range(26, 2))

    local offset = 28
    local param_count = buffer:range(25, 1):uint()

    if param_count > 0 then
        local param_tree = root:add(spw_proto, buffer(offset, param_count*4), "Parameters")
        while param_count > 0 do
            param_tree:add_le(pf_net2_dd_param, buffer:range(offset, 4))

            offset = offset + 4
            param_count = param_count - 1
        end
    end
    
    -- check if content exists
    if buffer:len() <= offset then
        return
    end

    data = buffer:range(offset)
    if buffer:range(24, 1):uint() == 1 then
        local obj = buffer:range(26, 2):le_uint()
        local resp_tree = root:add(spw_proto, data, "Response")

        off = 0
        while off + 8 <= data:len() do
            local data_type = data:range(off+3, 1):uint()

            local data_length = 0
            if data_type == 0x10 then
                data_length = data:range(off+8, 32):string():len()
            elseif data_type == 0x08 then
                data_length = 8*4
            elseif obj == 0x5400 then
                data_length = 8
            elseif data_type == 0x00 or data_type == 0x40 then
                data_length = 5*4
            end

            local val_tree = resp_tree:add(spw_proto, data:range(off, data_length+8), "Value")
            val_tree:add_le(pf_net2_dd_resp_class, data:range(off, 1))
            val_tree:add_le(pf_net2_dd_resp_code, data:range(off+1, 2))
            val_tree:add_le(pf_net2_dd_resp_type, data:range(off+3, 1))
            val_tree:add_le(pf_net2_dd_resp_timestamp, data:range(off+4, 4))

            if data_type == 0x10 then
                val_tree:add(pf_net2_dd_resp_value_s, data:range(off+8, data_length))
            elseif data_type == 0x08 then
                for i = 0,7,1 do
                    val = data:range(off+8+4*i, 3)

                    if val:le_uint() == 0xfffffe then
                        break
                    end
                    if data:range(off+8+4*i+3, 1):le_uint() == 1 then
                        val_tree:add_le(pf_net2_dd_resp_value, val)
                    end
                end
            elseif obj == 0x5400 then
                val_tree:add_le(pf_net2_dd_resp_value_l, data:range(off+8, 8))

            elseif data_type == 0x00 then
                for i = 0,4,1 do
                    val = data:range(off+8+4*i, 4)

                    if val:le_uint() == 0xffffffff then
                        break
                    end
                    val_tree:add_le(pf_net2_dd_resp_value, val)
                end
                
            elseif data_type == 0x40 then
                for i = 0,4,1 do
                    val = data:range(off+8+4*i, 4)

                    if val:le_uint() == -0x80000000 then
                        break
                    end
                    val_tree:add_le(pf_net2_dd_resp_value_i, val)
                end
            end

            off = off + 8 + data_length
        end
    else
        root:add(pf_spw_raw, data)
    end
end
